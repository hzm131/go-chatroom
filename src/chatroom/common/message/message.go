package message


const (
	LoginMesType = "LoginMes"
	LoginResMesType = "LoginResMes"
	RegisterMesType = "RegisterMes"
)

type Message struct {
	Type string `json:"type"`//消息类型
	Data string `json:"data"`//消息内容
}

type LoginMes struct {
	UserId int `json:"userId"`//用户id
	UserPwd string `json:"userPwd"`//用户密码
	UserName string `json:"userName"`//用户名
}

type LoginResMes struct {
	Code int `json:"code"`//状态码 500表示用户未注册 200成功
	Error string `json:"error"`//返回错误信息
}

type RegisterMes struct {
	//...
}