package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

func main() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("Client connect error ! " + err.Error())
		return
	}
	defer conn.Close()

	fmt.Println(conn.LocalAddr().String() + " :server connected!")
}

func onMessageReceived(conn *net.TCPConn) {

	reader := bufio.NewReader(conn)
	b := []byte(conn.LocalAddr().String() + " Say hello to Server... \n")
	conn.Write(b)

	for {
		msg, err := reader.ReadString('\n')
		fmt.Println("readString")
		fmt.Println(msg)
		if err != nil || err == io.EOF {
			fmt.Println(err)
			break
		}
		time.Sleep(2 * time.Second)

		fmt.Println("writing....")

		b := []byte(conn.LocalAddr().String() + " write data to Server... \n")
		_, err = conn.Write(b)

		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
