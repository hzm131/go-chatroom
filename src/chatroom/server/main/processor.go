package main

import (
	"chatroom/common/message"
	"chatroom/server/process"
	"chatroom/server/utils"
	"fmt"
	"io"
	"net"
)


//先创建一个Processor的结构体
type Processor struct {
	Conn net.Conn
}

//编写一个ServerProcessMes 函数
//功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func (this *Processor) ServerProcessMes(mes *message.Message)(err error){
	switch mes.Type {
	case message.LoginMesType:
		//	处理登录
		//创建一个UserProcess实例
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册
		up := &process2.UserProcess{
			Conn: this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		//创建一个UserProcess的实例，完成群发消息
		smsProcess := process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理。。。")
	}
	return
}


func (this *Processor) CreateProcess()(err error){
	//读客户端发送的信息
	for {
		//创建一个Transfer实例完成读包任务
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		mes,err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF{
				fmt.Println("客户端退出了，服务器也正常退出。。。")
				return err
			}
			fmt.Println("readPkg err=",err)
			return err
		}
		err = this.ServerProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}