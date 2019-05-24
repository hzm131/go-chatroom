package process

import (
	"chatroom/clinet/models"
	"chatroom/common/message"
	"fmt"
)

var onlineUser map[int]*message.User = make(map[int]*message.User,100)

var CurUser models.CurUser //在用户登录成功后完成对CurUser的初始化


func outPutOnlineUser(){
	fmt.Println("在线用户列表:")
	//遍历onlineUser
	for id,_ := range onlineUser{
		fmt.Println("用户id:\t",id)
	}
}

func upDataUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes){
	//适当优配
	user,ok := onlineUser[notifyUserStatusMes.UserId]
	if !ok {
		user = &message.User{
			UserId:notifyUserStatusMes.UserId,
		}
	}
	user.UserStatus = notifyUserStatusMes.UserStatus
	onlineUser[notifyUserStatusMes.UserId] = user

	outPutOnlineUser()
}