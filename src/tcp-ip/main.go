package main

import (
	"fmt"
	"io"
	"net"
)


func process(conn net.Conn){
	//循环的接受客户端发送的数据
	defer conn.Close() //关闭

	for {
		fmt.Printf("服务端在等待客户端%s 发送的信息\n",conn.RemoteAddr().String())
		//创建一个新的切片
		buf := make([]byte,1024)
		//conn.Read(buf)
		//1.等待客户端通过conn发送信息
		//2.如果客户端没有write[发送],那么协程就会一直堵塞在这里
		n,err := conn.Read(buf) // n是字节数
		if err == io.EOF { //客户端已退出
			// 超时，客户端退出 就会报错
			fmt.Println("客户端已退出")
			return
		}
		//显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[0:n])) // 这里代表真正读取到的数据 0 到 n
	}
}


func main(){
	// 使用tcp协议，在本地监听8888端口
	listen,err := net.Listen("tcp","0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err=",err)
		return
	}
	fmt.Printf("listen suc=%v",listen)
	defer listen.Close()
	//需要一直监听
	for {
		fmt.Println("等待客户端的连接...")
		conn,err := listen.Accept()
		if err != nil{
			fmt.Println("连接失败",err)
		}else{
			fmt.Println("连接成功",conn,conn.RemoteAddr().String())
		}

		//每个连接会用个新的协程去执行
		go process(conn)
	}

}
