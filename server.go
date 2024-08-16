package main

import (
	"fmt"
	"net"
	"sync"
)

/*
* 编译命令 go build -o server.exe .\main.go .\server.go
* 运行命令 .\server.exe
 */
type Server struct {
	Ip            string
	Port          int
	OnlineUserMap map[string]*User
	mapLock       sync.RWMutex
	MessageCh     chan string
}

// 创建一个服务端
func NewServer(ip string, port int) *Server {
	server := Server{
		Ip:            ip,
		Port:          port,
		OnlineUserMap: make(map[string]*User),
		MessageCh:     make(chan string),
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

	// 启动推送全部消息的goroutine
	go server.PushMsgToAllUser()

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

// 广播方法
func (server *Server) BroadCast(msg string) {
	server.MessageCh <- msg
}

// 将信息 推送给所有用户的方法
func (server *Server) PushMsgToAllUser() {
	for {
		msg := <-server.MessageCh

		// 推送消息给所有用户
		server.mapLock.Lock()
		for _, cli := range server.OnlineUserMap {
			cli.Ch <- msg
		}
		server.mapLock.Unlock()
	}

}

// 处理当前业务的方法
func (server *Server) Handler(conn net.Conn) {
	//fmt.Println("connection establish successfully!")

	// 1.创建用户实例
	userPtr := NewUser(conn)

	// 2.将用户实例存入onlineMap
	server.mapLock.Lock()
	server.OnlineUserMap[userPtr.Name] = userPtr
	server.mapLock.Unlock()

	// 3. 广播用户上线信息
	server.BroadCast(userPtr.Name + ": Online!\n")

}
