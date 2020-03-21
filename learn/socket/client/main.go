package main

import (
	"fmt"
	"net"
)

func main() {
	//客户端，获取服务端地址和端口
	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}
	defer conn.Close()

	conn.Write([]byte("helloworld"))

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}

	fmt.Println("服务器回发:", string(buf[:n]))
}
