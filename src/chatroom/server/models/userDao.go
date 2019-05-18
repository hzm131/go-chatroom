package models

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)


//我们在服务器启动后，就初始化一个UserDao
//把他做成全局变量，在需要和redis操作时，直接使用即可
var (
	MyUserDao *UserDao
)

//定义一个UserDao结构体，完成对User结构体的各种操作
type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao){
	userDao = &UserDao{
		pool:pool,
	}
	return
}

//1.根据一个用户id，返回一个User实例+err
func (this *UserDao) getUserById(conn redis.Conn,Id int)(user *User,err error){
	//通过给定的id去redis查询这个用户
	res,err := redis.String(conn.Do("hget","users",Id))
	if err != nil{
		//错误
		if err == redis.ErrNil{ //表示在user hash中没有找到对应的id
			err = ERROR_USER_NOTEXISTS
			fmt.Println("用户不存在")
		}
		fmt.Println("内部错误")
		return
	}
	user = &User{}
	//用户存在res是json字符串，所以需要将res反序列化成user实例
	err = json.Unmarshal([]byte(res),user)
	if err != nil{
		fmt.Println("json.Unmarshal err=",err)
		return
	}

	return
}

//完成登录校验 Login
//1.Login 完成对用户的验证
//2.如果用户的id和pwd都正确，则返回一个user实例
//3.如果id或pwd有误，则返回对应的错误信息
func (this *UserDao) Login(userId int,userPwd string)(user *User,err error){
	//先从UserDao的连接池取一根连接
	conn := this.pool.Get()
	defer conn.Close()

	user,err = this.getUserById(conn,userId)
	if err != nil{
		return
	}
	//到这一步说明用户获取到了，然后就需要将从redis中取出的用户与用户发送的用户数据对比
	if user.UserPwd != userPwd{
		err = ERROR_USER_PWD
		fmt.Println("密码错误")
		return
	}
	return
}
