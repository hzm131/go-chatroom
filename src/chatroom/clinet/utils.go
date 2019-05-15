package main

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func readPkg(conn net.Conn)(mes message.Message,err error){
	fmt.Println("读取客户端发送的数据...")
	buf := make([]byte,8096)
	_,err = conn.Read(buf[:4])
	if err!= nil{
		//err = errors.New("conn.Read header err")
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[:4])
	n,err := conn.Read(buf[:pkgLen]) //从套接字中读pkgLen个字节到buf中
	if n != int(pkgLen) || err != nil{
		//err = errors.New("conn.Read body err")
		return
	}

	//把pkgLen反序列化成->message.Message
	err = json.Unmarshal(buf[:pkgLen],&mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=",err)
		return
	}
	return
}


func WritePkg(conn net.Conn,data []byte)(err error){
	//先发送一个长度给对方
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
	//发送data本身
	n,err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn,Write(bytes) fail",err)
		return
	}
	return
}

//编写一个ServerProcessLogin函数，专门处理登录请求
func ServerProcessLogin(conn net.Conn,mes *message.Message)(err error){
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
	err = WritePkg(conn,data)
	return
}
