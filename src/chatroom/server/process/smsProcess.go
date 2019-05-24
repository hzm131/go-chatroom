package process2

import (
	"chatroom/clinet/utils"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"net"
)

type SmsProcess struct {

}

//写方法转发消息
func(this *SmsProcess)SendGroupMes(mes *message.Message){
	//遍历服务器端的map onlineUsers map[int]*UserProcess
	//将消息转发出去

	//取出mes中的内容
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil {
		fmt.Println("反序列化失败",err)
		return
	}

	data,err := json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes序列化失败",err)
		return
	}

	for id,up := range userMgr.onlineUsers{
		//这里还需要过滤掉自己，不要发给自己
		if id == smsMes.UserId{
			continue
		}
		this.SendMesToEachOnlineUsers(data,up.Conn)
	}
}


func(this *SmsProcess)SendMesToEachOnlineUsers(data []byte,conn net.Conn){
	//创建一个Transfer,实例，发送data
	tf := utils.Transfer{
		Conn:conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发消息失败",err)
	}
}