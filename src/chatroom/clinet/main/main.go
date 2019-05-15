package main

import (
	"chatroom/clinet/process"
	"fmt"
	"os"
	"time"
)


//定义两个变量:一个表示用户id，一个表示用户密码
var userId int
var userPwd string

func main(){
	// 接受用户选择
	var key int

	for true {
		fmt.Println("------欢迎登陆多人聊天系统-----")
		fmt.Println("\t\t\t 1.登陆聊天室")
		fmt.Println("\t\t\t 2.注册用户")
		fmt.Println("\t\t\t 3.退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")
		fmt.Scanf("%d\n",&key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			fmt.Println("请输入用户的id:")
			fmt.Scanln(&userId)
			fmt.Println("请输入用户的密码:")
			fmt.Scanln(&userPwd)
			// 完成登录
			// 1.创建一个UserProcess的实例
			up := &process.UserProcess{}
			up.Login(userId,userPwd)
		case 2:
			fmt.Println("注册用户")
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误,请重新输入")
		}
	}
	//根据用户的输入，显示新的提示信息

	time.Sleep(time.Second * 5)
}
