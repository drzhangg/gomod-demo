package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"time"
)

type Message struct {
	Username string
	Message  string
}

type User struct {
	Username string
}

type Datas struct {
	Messages []Message
	Users    []User
}

// 全局信息
var datas Datas
var users map[*websocket.Conn]string

func main() {
	fmt.Println("启动时间: ", time.Now())

	// 初始化数据
	datas = Datas{}
	users = make(map[*websocket.Conn]string)

	// 渲染页面
	http.HandleFunc("/", index)

	// 监听socket方法
	http.Handle("/webSocket", websocket.Handler(webSocket))

	// 监听8080端口
	http.ListenAndServe(":8889", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func webSocket(ws *websocket.Conn) {
	var message Message
	var data string
	for {
		// 接收数据
		err := websocket.Message.Receive(ws, &data)
		if err != nil {
			// 移除出错的连接
			delete(users, ws)
			fmt.Println("连接异常")
			break
		}
		// 解析信息
		err = json.Unmarshal([]byte(data), &message)
		if err != nil {
			fmt.Println("解析数据异常")
		}

		// 添加新用户到map中,已经存在的用户不必添加
		if _, ok := users[ws]; !ok {
			users[ws] = message.Username
			// 添加用户到全局信息
			datas.Users = append(datas.Users, User{Username: message.Username})
		}
		// 添加聊天记录到全局信息
		datas.Messages = append(datas.Messages, message)

		// 通过webSocket将当前信息分发
		for key := range users {
			err := websocket.Message.Send(key, data)
			if err != nil {
				// 移除出错的连接
				delete(users, key)
				fmt.Println("发送出错: " + err.Error())
				break
			}
		}
	}
}
