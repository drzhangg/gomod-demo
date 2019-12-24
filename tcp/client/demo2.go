package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

var host = flag.String("host", "", "host")
var port = flag.String("port", "9999", "port")

type Msg struct {
	Data string `json:"data"`
	Type int    `json:"type"`
}

type Resp struct {
	Data   string `json:"data"`
	Status int    `json:"status"`
}

func main() {
	flag.Parse()

	var listen net.Listener
	var err error
	listen, err = net.Listen("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("listener err:", err)
		os.Exit(1)
	}
	defer listen.Close()

	fmt.Println("listening on ", *host, ":", *port)

	for {
		conn, err := listen.Accept()
		if err != nil || err == io.EOF {
			fmt.Println("accept err:", err)
			os.Exit(1)
		}
		fmt.Printf("message %s -> %s\n", conn.RemoteAddr().String(), conn.LocalAddr())

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	ip := conn.RemoteAddr().String()
	defer func() {
		fmt.Println("disconnect:" + ip)
		conn.Close()
	}()

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	for {
		b, _, err := reader.ReadLine()
		if err != nil {
			return
		}

		var msg Msg
		json.Unmarshal(b, &msg)
		fmt.Println("GET==>data ", msg.Data, " type:", msg.Type)

		resp := Resp{
			Data:   time.Now().String(),
			Status: 200,
		}

		datas, _ := json.Marshal(resp)
		writer.Write(datas)
		writer.Write([]byte("\n"))
		writer.Flush()
	}
	fmt.Println("done")
}
