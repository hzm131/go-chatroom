package message

const (
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
	RegisterResMesType = "RegisterResMes"
	NotifyUserStatusMesType = "NotifyUserStatusMes"
	SmsMesType = "SmsMes"
)

//定义几个用户状态的常量
const(
	UserOnLine = iota //在线
	UserOffLline
	UserBusyStatus
)



type Message struct {
	Type string `json:"type"`//消息类型
	Data string `json:"data"`//消息内容
}

//登录
type LoginMes struct {
	UserId int `json:"userId"`//用户id
	UserPwd string `json:"userPwd"`//用户密码
	UserName string `json:"userName"`//用户名
}
//响应
type LoginResMes struct {
	Code int `json:"code"`//状态码 500表示用户未注册 200成功
	UsersId []int //增加字段，保存用户ud的切片
	Error string `json:"error"`//返回错误信息
}

//注册
type RegisterMes struct {
	User User `json:"user"`
}
//响应
type RegisterResMes struct {
	Code int `json:"code"` //返回状态码 400表示该用户已占用
	Error string `json:"error"`
}

//为了配合服务器推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"` //用户id
	UserStatus int `json:"userStatus"` //用户状态
}


//增加一个SmsMes 发送的
type SmsMes struct {
	Content string `json:"content"` //内容
	User
}

