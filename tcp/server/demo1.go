package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	//定义一个tcp断点
	var tcpAddr *net.TCPAddr

	//通过ResolveTCPAddr实例一个具体的tcp断点
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

	//打开一个tcp断点监听
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()
	fmt.Println("server ready to read...")

	//循环接收客户端的连接，创建一个协程去具体处理连接
	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("A client connect : ", tcpConn.RemoteAddr().String())
		go tcpPipe(tcpConn)
	}
}

//具体处理连接过程方法
func tcpPipe(conn *net.TCPConn) {
	//tcp连接的地址
	ipStr := conn.RemoteAddr().String()

	defer func() {
		fmt.Println("disconnected: ", ipStr)
		conn.Close()
	}()

	//获取一个连接的reader读取流
	reader := bufio.NewReader(conn)
	i := 0
	//接收并返回消息
	for {
		message, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		fmt.Println(string(message))

		time.Sleep(time.Second * 3)

		msg := time.Now().String() + conn.RemoteAddr().String() + " Server Say hello! \n"

		b := []byte(msg)

		conn.Write(b)

		i++

		if i > 10 {
			break
		}
	}
}
