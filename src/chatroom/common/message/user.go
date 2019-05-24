package message
//定义一个用户的结构体
//为了保证user序列化或反序列化成功，用户信息的json字符串和对应的tag名字一致
type User struct {
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
	UserStatus int `json:"userStatus"` //用户状态
	//可以扩展
	Sex string `json:"sex"` //性别
}