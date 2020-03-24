package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:8003")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close()

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Read err:", err)
		return
	}
	fmt.Println("服务端消息是：", string(buf[:n]))

	_, err = conn.Write([]byte("收到客户端消息"))
	if err != nil {
		fmt.Println("Read err:", err)
		return
	}
}
