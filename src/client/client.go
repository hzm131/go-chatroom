package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// 客户端
func main(){
	conn,err := net.Dial("tcp","127.0.0.1:8888")
	if err != nil{
		fmt.Println("客户端连接服务器失败",err)
		return
	}
	//功能1：客户端可以发送单行数据，然后退出
	reader := bufio.NewReader(os.Stdin)  //Stdin代表标准输入->终端


	for{
		// 从终端读取一行用户输入，并准备发送给服务器
		line,err := reader.ReadString('\n')
		if err != nil{
			fmt.Println("读取失败 err=",err)
		}

		line = strings.Trim(line," \r\n")
		if line == "exit"{
			fmt.Println("客户端退出。。。")
			break
		}
		//再将 line 发送给服务器
		_,err = conn.Write([]byte(line + "\n")) //n代表返回了多少个字节
		if err != nil{
			fmt.Println("写入失败",err)
		}

	}
}
