package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)
// mysql 连接
//使用第三方开源的mysql库: github.com/go-sql-driver/mysql （mysql驱动）
// github.com/jmoiron/sqlx （基于mysql驱动的封装）
// database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")

type Person struct {
	UserId int `db:"user_id"`
	Username string `db:"username"`
	Sex string `db:"sex"`
	Email string `db:"email"`
}
var db *sqlx.DB

func init()  {
	open, err := sqlx.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		fmt.Println("open mysql failed, error:", err)
		return
	}
	db = open
}

func insert(username,sex,email string)  {
	result, err := db.Exec("insert into person(username, sex, email) values (?,?,?)", username, sex, email)
	if err != nil {
		fmt.Println("mysql insert data failed, error:", err)
		return
	}
	i, err := result.LastInsertId()
	if err != nil {
		fmt.Println("get id failed, error:", err)
		return
	}
	fmt.Println("insert success id:",i)
}

func getData (id int)  {
	var person  []Person
	err := db.Select(&person, "select user_id,username,sex,email from person where user_id = ?", id)
	if err != nil {
		fmt.Println("mysql select failed, error:", err)
		return
	}
	fmt.Println("select success data;", person)
}

func update(username, email string ,id int)  {
	result, err := db.Exec("update person set username = ?,email = ?  where user_id = ?", username, email, id)
	if err != nil {
		fmt.Println("mysql insert data failed, error:", err)
		return
	}
	i, err := result.RowsAffected()
	if err != nil {
		fmt.Println("get id failed, error:", err)
		return
	}
	fmt.Println("Affect success id:",i)
}

func del(id int)  {
	result, err := db.Exec("delete from proson where user_d = ?", id)
	if err != nil {
		fmt.Println("mysql insert data failed, error:", err)
		return
	}
	i, err := result.RowsAffected()
	if err != nil {
		fmt.Println("get id failed, error:", err)
		return
	}
	fmt.Println("Affect success id:",i)
}

/*
mysql事务特性：
    1) 原子性
    2) 一致性
    3) 隔离性
    4) 持久性
golang MySQL事务应用：
    1） import (“github.com/jmoiron/sqlx")
    2)  Db.Begin()        开始事务
    3)  Db.Commit()        提交事务
    4)  Db.Rollback()     回滚事务
*/
func acid()  {
	conn, err := db.Begin()
	if err != nil {
		fmt.Println("begin failed, error:", err)
		return
	}
	r, err := conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)

	r, err = conn.Exec("insert into person(username, sex, email)values(?, ?, ?)", "stu001", "man", "stu01@qq.com")
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	id, err = r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed, ", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert succ:", id)

	conn.Commit()
}

func main() {
	insert("stu01", "man", "stu01@admin.com")
}
