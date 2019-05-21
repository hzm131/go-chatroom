package process2

import (
	"chatroom/common/message"
	"chatroom/server/models"
	"chatroom/server/utils"
	"encoding/json"
	"fmt"
	"net"
)


type UserProcess struct {
	Conn net.Conn
	//增加一个字段，表示该Conn是哪个用户的
	UserId int
}

func (this *UserProcess)ServerProcessRegister(mes *message.Message)(err error){
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data),&registerMes)
	if err != nil{
		fmt.Println("json.Unmarshal fail err=",err)
		return
	}
	//1.先声明一个resMes
	var resMes message.Message
	resMes.Type = message.RegisterMesType

	//2.再声明一个LoginResMes，并完成赋值
	var registerResMes message.RegisterResMes
	//到数据库中去完成注册
	err = models.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == models.ERROR_USER_EXISTS{
			registerResMes.Code = 505
			registerResMes.Error = err.Error()
		}else{
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	}else{
		registerResMes.Code = 200
	}
	data,err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail",err)
		return
	}
	//4.将data赋值给resMes
	resMes.Data = string(data)

	//5.对resMes 进行序列化，准备发送
	data,err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail",err)
		return
	}
	//6.发送data 我们将其封装到writePkg函数
	//因为使用了分层模式，先创建一个Transfer实例，然后读取
	tf := utils.Transfer{
		Conn: this.Conn,

	}
	err = tf.WritePkg(data)
	return
}

//编写一个ServerProcessLogin函数，专门处理登录请求
func (this *UserProcess)ServerProcessLogin(mes *message.Message)(err error){
	//1.先从mes中取出mes.Data，并直接反序列化LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil{
		fmt.Println("json.Unmarshal fail err=",err)
		return
	}

	//1.先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//2.再声明一个LoginResMes，并完成赋值
	var loginResMes message.LoginResMes

	//使用model.MyuserDao 到redis去验证
	user,err := models.MyUserDao.Login(loginMes.UserId,loginMes.UserPwd)
	if err != nil{

		if err == models.ERROR_USER_NOTEXISTS{
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		}else if err == models.ERROR_USER_PWD{
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		}else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}
		//这里我们先测试成功，然后我们可以根据返回具体错误信息
	}else{
		loginResMes.Code = 200
		//这里，因为用户登录成功，我们就把该登录成功的用户放入userMgr中
		//将登录成功的用户的userId赋给this
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		//将当前在线用户的id放入到loginResMes.UsersId
		//便利UserMgr.onlineUsers
		for id,_ := range userMgr.onlineUsers{
			loginResMes.UsersId = append(loginResMes.UsersId,id)
		}
		fmt.Println(user,"登录成功")
	}
	//3.将loginResMes序列化
	data,err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail",err)
		return
	}
	//4.将data赋值给resMes
	resMes.Data = string(data)

	//5.对resMes 进行序列化，准备发送
	data,err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail",err)
		return
	}
	//6.发送data 我们将其封装到writePkg函数
	//因为使用了分层模式，先创建一个Transfer实例，然后读取
	tf := utils.Transfer{
		Conn: this.Conn,

	}
	err = tf.WritePkg(data)
	return
}
