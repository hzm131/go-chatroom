package process

import (
	"chatroom/clinet/utils"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

//显示登录成功后的函数

func ShowMenu(){
	fmt.Println("------恭喜XXX登录成功------")
	fmt.Println("------1.显示在线用户列表------")
	fmt.Println("------2.发送消息------")
	fmt.Println("------3.信息列表------")
	fmt.Println("------4.退出系统------")
	fmt.Println("请选择1-4:")
	var key int
	var content string
	//因为我们总会使用到smsProcess的实例，因此将其定义在switch的外部，避免重复创建实例带来的消耗
	smsProcess := SmsProcess{}
	fmt.Scanf("%d\n",&key)
	switch key {
		case 1:
			fmt.Println("显示在线用户列表")
			outPutOnlineUser()
		case 2:
			fmt.Println("你想对大家说点什么:")
			fmt.Scanf("%s\n",&content)
			smsProcess.SendGroupMes(content)
		case 3:
			fmt.Println("信息列表")
		case 4:
			fmt.Println("你选择退出了系统。。。")
			os.Exit(0)
		default:
			fmt.Println("你输入的选项不正确。。")
	}
}

//和服务器端保持通讯
func ServerProcessMes(conn net.Conn){
	// 创建一个Transfer实例，不停的读取服务器发送的消息
	tf := &utils.Transfer{
		Conn:conn,
	}
	for {
		//fmt.Printf("客户端%s正在读取服务器发送的消息")
		mes ,err := tf.ReadPkg()
		if err != nil{
			fmt.Println("tf.ReadPkg err=",err)
			return
		}
		//如果读取到了消息,又是下一步处理逻辑
		//fmt.Printf("mes=%v\n",mes)
		switch mes.Type {
			case message.NotifyUserStatusMesType: //有人上线了
			//处理
			//1.取出NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			json.Unmarshal([]byte(mes.Data),&notifyUserStatusMes)
			//2.把这个用户信息，状态保存到客户端map[int]User中
			upDataUserStatus(&notifyUserStatusMes)

			case message.SmsMesType:  //有人群发消息了
				outPutGroupMes(&mes)
			default:
				fmt.Println("服务器端返回未知的消息类型")
		}
	}
}