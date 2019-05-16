package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

//定义一个全局的pool
var pool *redis.Pool

func initPool(address string,maxIdle,maxActive int,idleTimeout time.Duration){
	pool = &redis.Pool{
		MaxIdle:maxIdle,  //最大空闲连接数
		MaxActive:maxActive, //表示数据库最大的连接数， 0表示没有限制
		IdleTimeout:idleTimeout, //最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp",address)
		},
	}
}
