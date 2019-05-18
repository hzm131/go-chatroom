package main

import (
	"chatroom/server/models"
	"fmt"
	"net"
	"time"
)

/*
1.先完成指定用户的验证，用户id是100，用户密码是123456，其他用户不能登录
2.完成客户端发送消息本身，服务端可以正常接收到消息，并根据客户端发送的消息（LoginMes）,判断用户的合法性，并返回相应的LoginResMes
	思路分析：
		1.让客户端发送消息本身
		2.服务器接收到消息，然后反序列化成对应的消息结构体
		3.服务器根据反序列化的消息，判断是否登录用户是合法用户，返回LoginResMes
		4.客户端解析返回的LoginResMes显示对应的界面
		5.需要做一些函数的封装



分层：
	1.先把分析出来的文件，创建好，然后方法相应的文件包
	2.根据各个文件完成的任务不通，将main.go文件中的代码剥离到对应的文件中
*/


//处理和客户端的通信
func process(conn net.Conn){
	//这里需要延时关闭
	defer conn.Close()
	//调用总控
	process := &Processor{
		Conn:conn,
	}
	err := process.CreateProcess()
	if err != nil{
		fmt.Println("客户端和服务器端通讯的协程出错了。。。",err)
		return
	}
}

func init(){
	//当服务器启动时，我们就去初始化redis的连接池
	initPool("127.0.0.1:6379",16,0,300*time.Second)
	initUserDao()
}

//这里编写一个函数，完成UserDao初始化任务
func initUserDao(){
	models.MyUserDao = models.NewUserDao(pool)
}

func main()  {
	listen,err := net.Listen("tcp","127.0.0.1:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("监听出错")
		return
	}
	//监听成功就等待客户端连接服务器
	for {
		fmt.Println("等待客户端来连接服务器")
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err =",err)
		}
		//一旦连接成功则启动一个协程和客户端保持通讯
		go process(conn)
	}
}