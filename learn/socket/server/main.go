package main

import (
	"fmt"
	"net"
)

func main() {
	//监听端口，生成一个socket
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	fmt.Println("服务端等待连接。。。。")

	//这里当没有服务端进行连接时，会发生阻塞    也是生成一个socket
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err:", err)
		return
	}
	defer conn.Close()

	fmt.Println("服务端连接成功！")

	bytes := []byte{}
	n, err := conn.Read(bytes)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}

	fmt.Println("服务端读到数据为：", string(bytes[:n]))

}
