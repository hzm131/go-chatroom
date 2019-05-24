package process

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
)

func outPutGroupMes(mes *message.Message){ //这个地方mes一定是SmsMes
	//显示即可
	//1.反序列化mes.Data
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data),&smsMes)
	if err != nil{
		fmt.Println("outPutGroupMes反序列化失败")
		return
	}
	//显示信息
	info := fmt.Sprintf("用户id:\t%d 对大家说:\t%s",smsMes.UserId,smsMes.Content)
	fmt.Println(info)
	fmt.Println()
}
