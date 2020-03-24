package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8003")
	if err != nil {
		fmt.Println("ResolveUDPAddr err:", err)
		return
	}
	fmt.Println("服务端获取地址成功")
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("ListenUDP err:", err)
		return
	}
	defer udpConn.Close()

	fmt.Println("服务端连接成功")

	buf := make([]byte, 4096)
	//读取数据
	n, addr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("ReadFromUDP err:", err)
		return
	}
	fmt.Printf("服务器读到 %v 的数据：%s\n", addr, string(buf[:n]))

	//向客户端回写数据
	nowTime := time.Now().String()
	_, err = udpConn.WriteToUDP([]byte(nowTime), addr)
	if err != nil {
		fmt.Println("WriteToUDP err:", err)
		return
	}

}
