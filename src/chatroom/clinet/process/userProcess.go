package process

import (
	"chatroom/clinet/utils"
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {

}

func (this *UserProcess)Register(userId int, userPwd string,userName string)(err error){
	//1.连接到服务器
	conn,err := net.Dial("tcp","127.0.0.1:8889")
	if err != nil{
		fmt.Println("连接失败，不完了")
		return
	}
	defer conn.Close()
	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.RegisterMesType
	//3.创建LoginMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName
	//4.将LoginMes进行序列化
	data,err := json.Marshal(registerMes)
	if err != nil{
		fmt.Println("json.Marshal err=",err)
		return
	}
	//5.就可以放入mes.Data中
	mes.Data = string(data)
	//6.将mes进行序列化，这是它才是可以发送的结构体
	data,err = json.Marshal(mes)
	if err != nil{
		fmt.Println("json.Marshal err=",err)
		return
	}
	tf := &utils.Transfer{
		Conn: conn,
	}
	//发送data给服务器端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 err=",err)
	}
	//读取服务器返回的信息
	mes,err = tf.ReadPkg() //mes就是RegisterResMes
	if err != nil{
		fmt.Println("readPkg err=",err)
		return
	}
	//将mes的Data部分反序列化成RegisterResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data),&registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，请重新登录")
		//os.Exit(0)
	}else {
		fmt.Println(registerResMes.Error)
		//os.Exit(0)
	}
	return
}

// 关联用户登录的方法
func (this *UserProcess) Login(userId int, userPwd string)(err error){
	//1.连接到服务器
	conn,err := net.Dial("tcp","127.0.0.1:8889")
	if err != nil{
		fmt.Println("连接失败，不完了")
		return
	}
	defer conn.Close()
	//2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType
	//3.创建LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4.将LoginMes进行序列化
	data,err := json.Marshal(loginMes)
	if err != nil{
		fmt.Println("json.Marshal err=",err)
		return
	}
	//5.就可以放入mes.Data中
	mes.Data = string(data)
	//6.将mes进行序列化，这是它才是可以发送的结构体
	data,err = json.Marshal(mes)
	if err != nil{
		fmt.Println("json.Marshal err=",err)
		return
	}
	//7.这个时候data就是我们要发送消息了
	//7.1 先把data的长度发送给服务器
	//先获取到data的长度->转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen)

	//发送长度
	n,err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn,Write(bytes) fail",err)
		return
	}
	//fmt.Printf("客户端，发送消息的长度=%d 内容=%s",len(data),string(data))
	_,err = conn.Write(data)
	if  err != nil {
		fmt.Println("conn,Write(bytes) fail",err)
		return
	}
	//处理服务器返回的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes,err = tf.ReadPkg() //mes就是
	if err != nil{
		fmt.Println("readPkg err=",err)
		return
	}

	//将mes的Data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data),&loginResMes)
	if loginResMes.Code == 200 {
		//登录成功后显示当前在线用户列表,遍历loginResMes.UsersId
		fmt.Println("当前在线用户列表如下:")
		for _,v := range loginResMes.UsersId{
			//如果要求不显示自己在线
			if v == userId{
				continue
			}
			fmt.Println("用户id:\t",v)
			//完成客户端的onlineUser的初始化
			user := &message.User{
				UserId:v,
				UserStatus:message.UserOnLine,
			}
			onlineUser[v] = user
		}
		fmt.Print("\n\n")
		//这里我们还需要在客户端启动一个协程
		//该协程保持和服务器端的通讯，如果服务器有数据推送给客户端
		//则接受并显示在客户端的终端
		go ServerProcessMes(conn)
		// 1.显示登录成功后的菜单[循环]
		for{
			ShowMenu()
		}
	}else {
		fmt.Println(loginResMes.Error)
	}
	return
}