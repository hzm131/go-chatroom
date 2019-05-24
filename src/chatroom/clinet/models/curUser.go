package models

import (
	"chatroom/common/message"
	"net"
)

//因为在客户端我们在很多地方会使用到CurUser，所用将其作为全局的
type CurUser struct {
	Conn net.Conn
	message.User
}