package process

import "fmt"

//显示登录成功后的函数

func ShowMenu(){
	fmt.Println("------恭喜XXX登录成功------")
	fmt.Println("------1.显示在线用户列表------")
	fmt.Println("------2.发送消息------")
	fmt.Println("------3.信息列表------")
	fmt.Println("------4.退出系统------")
	fmt.Println("------请选择1-4------")
	var key int
	fmt.Scanf("%d\n",&key)
	switch key {
		case 1:
			fmt.Println()
		case 2:
			fmt.Println()
		case 3:
			fmt.Println()
		case 4:
			fmt.Println()
		default:
			fmt.Println()
	}
}
