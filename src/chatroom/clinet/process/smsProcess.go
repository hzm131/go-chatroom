package process

import (
	"chatroom/clinet/utils"
	"chatroom/common/message"
	"encoding/json"
	"fmt"
)

type SmsProcess struct {

}

//发送群聊的消息
func(this *SmsProcess)SendGroupMes(content string)(err error){
	//1.创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//2.创建一个SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus

	//序列化smsMes
	data,err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes序列化失败",err)
		return
	}
	mes.Data = string(data)
	//4.对mes再次序列化
	data,err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes序列化失败",err)
		return
	}
	//5.将序列化后的mes发送给服务器
	tf := utils.Transfer{
		Conn:CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes发送信息失败",err)
		return
	}
	return
}