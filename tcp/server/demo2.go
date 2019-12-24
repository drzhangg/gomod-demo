package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
)

var host = flag.String("host", "localhost", "host")
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

	conn, err := net.Dial("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("connect err :", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("connecting to " + *host + ":" + *port)

	var wg sync.WaitGroup

	wg.Add(2)
	go handleWrite(conn, &wg)
	go handleRead(conn, &wg)
	wg.Wait()
}

func handleWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 10; i > 0; i-- {
		d := "hello" + strconv.Itoa(i)
		msg := Msg{
			Data: d,
			Type: i,
		}

		datas, _ := json.Marshal(msg)
		writer := bufio.NewWriter(conn)
		writer.Write(datas)
		writer.Write([]byte("\n"))
		writer.Flush()
	}
	fmt.Println("write done")
}

func handleRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	reader := bufio.NewReader(conn)
	for i := 0; i < 10; i++ {
		b, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println("read error", err)
			return
		}

		var resp Resp
		json.Unmarshal(b, &resp)
		fmt.Println("status", resp.Status, " content:", resp.Data)
	}
	fmt.Println("read done")
}
