package main

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
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

*/
func readPkg(conn net.Conn)(mes message.Message,err error){
	fmt.Println("读取客户端发送的数据...")
	buf := make([]byte,8096)
	// conn.Read只有在未关闭的情况下，才会阻塞
	//如果客户端关系了conn，就不会阻塞
	_,err = conn.Read(buf[:4])
	if err!= nil{
		 //err = errors.New("conn.Read header err")
		return
	}
	// 根据buf[:4]读到的长度转成uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])
	n,err := conn.Read(buf[:pkgLen]) //从套接字中读pkgLen个字节到buf中
	if n != int(pkgLen) || err != nil{
		//err = errors.New("conn.Read body err")
		return
	}

	//把pkgLen反序列化成->message.Message
	err = json.Unmarshal(buf[:pkgLen],&mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=",err)
		return
	}
	return
}

//编写一个ServerProcessLogin函数，专门处理登录请求
func ServerProcessLogin(conn net.Conn,mes *message.Message)(err error){
	//1.先从mes中取出mes.Data，并直接反序列化LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil{
		fmt.Println("json.Unmarshal fail err=",err)
		return
	}

	//1.先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//2.再声明一个LoginResMes，并完成赋值
	var loginResMes message.LoginResMes

	//如果用户的id = 100 密码 = 123456，认为是合法的，否则不合法
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456"{
		//合法
		loginResMes.Code = 200

	}else{
		//不合法
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在，请注册再使用"
	}
	//3.将loginResMes序列化
	data,err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail",err)
		return
	}
	//4.将data赋值给resMes
	resMes.Data = string(data)

	//5.对resMes 进行序列化，准备发送
	data,err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail",err)
		return
	}
	//6.发送data 我们将其封装到writePkg函数
	err = WritePkg(conn,data)
	return
}

func WritePkg(conn net.Conn,data []byte)(err error){
	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen)

	//发送长度
	n,err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn,Write(bytes) fail",err)
		return
	}
	//发送data本身
	n,err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn,Write(bytes) fail",err)
		return
	}
	return
}

//编写一个ServerProcessMes 函数
//功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func ServerProcessMes(conn net.Conn,mes *message.Message)(err error){
	switch mes.Type {
		case message.LoginMesType:
			//	处理登录
			err = ServerProcessLogin(conn,mes)
		case message.RegisterMesType:
			//处理注册
		default:
			fmt.Println("消息类型不存在，无法处理。。。")
	}
	return
}



//处理和客户端的通信
func process(conn net.Conn){
	//这里需要延时关闭
	defer conn.Close()

	//读客户端发送的信息
	for {
		mes,err := readPkg(conn)
		if err != nil {
			if err == io.EOF{
				fmt.Println("客户端退出了，服务器也正常退出。。。")
				return
			}
			fmt.Println("readPkg err=",err)
			return
		}
		err = ServerProcessMes(conn,&mes)
		if err != nil {
			return
		}
	}
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