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
		loginResMes.Code = 500
		loginResMes.Error = "还用户不存在，请注册再使用"
		//这里我们先测试成功，然后我们可以根据返回具体错误信息
	}else{
		loginResMes.Code = 200
		fmt.Println(user,"登录成功")
	}

	//如果用户的id = 100 密码 = 123456，认为是合法的，否则不合法
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456"{
		//合法
		loginResMes.Code = 200

	}else{
		//不合法
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在，请注册再使用"
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