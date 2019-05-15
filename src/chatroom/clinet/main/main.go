package main

import (
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
	// 判断是否还继续显示菜单
	var loop = true

	for loop {
		fmt.Println("------欢迎登陆多人聊天系统-----")
		fmt.Println("\t\t\t 1.登陆聊天室")
		fmt.Println("\t\t\t 2.注册用户")
		fmt.Println("\t\t\t 3.退出系统")
		fmt.Println("\t\t\t 请选择(1-3):")
		fmt.Scanf("%d\n",&key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("你的输入有误,请重新输入")
		}
	}
	//根据用户的输入，显示新的提示信息
	if key == 1 {
		//说明用户要登陆
		fmt.Println("请输入用户的id:")
		fmt.Scanln(&userId)
		fmt.Println("请输入用户的密码:")
		fmt.Scanln(&userPwd)
		//把登录写到另一个文件
		//这里需要从新调用
		login(userId,userPwd)
	}else if key == 2 {
		fmt.Println("进行用户注册的逻辑")
	}
	time.Sleep(time.Second * 5)
}
