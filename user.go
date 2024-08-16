package main

import "net"

type User struct {
	Name string
	Addr string
	Ch   chan string
	conn net.Conn
}

// 新建一个用户
func NewUser(conn net.Conn) *User {
	userAddr := conn.RemoteAddr().String()
	user := User{
		Name: userAddr,
		Addr: userAddr,
		Ch:   make(chan string),
		conn: conn,
	}
	// 开启用户信息监听的gorouting
	go user.ListenMsg()
	return &user
}

func (user *User) ListenMsg() {
	for {
		msg := <-user.Ch
		user.conn.Write([]byte(msg + "\n"))
	}
}
