package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool //创建redis连接池

// //实例化一个连接池
func init()  {
	pool =  &redis.Pool{
		MaxIdle: 16,  //最初的连接数量
		MaxActive:0, //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300 , ////连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) {
			return  redis.Dial("tcp", "localhost:6379")
		},
	}
}

func main() {

	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("redis connection failed, error:", err)
		return
	}
	fmt.Println("redis connection success")
	defer conn.Close()
	_, err = conn.Do("set", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	i, err := redis.Int(conn.Do("get", "abc"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(i)
}
