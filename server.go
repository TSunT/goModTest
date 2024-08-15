package main

import (
	"fmt"
	"net"
)

/*
* 编译命令 go build -o server.exe .\main.go .\server.go
* 运行命令 .\server.exe
 */

type Server struct {
	Ip   string
	Port int
}

// 创建一个服务端
func goNewServer(ip string, port int) *Server {
	server := Server{
		Ip:   ip,
		Port: port,
	}
	return &server
}

// 开启一个服务端
func (server *Server) Start() {
	// 监听
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server.Ip, server.Port))
	if err != nil {
		fmt.Println("start listen failed: ", err)
		return
	} else {
		fmt.Println("start listen succeed! ")
	}
	defer listener.Close()
	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept failed: ", err)
			continue
		}
		// 开启业务goroutine
		go server.Handler(conn)
	}
}

// 处理当前业务的方法
func (server *Server) Handler(conn net.Conn) {
	fmt.Println("connection establish successfully!")
}
