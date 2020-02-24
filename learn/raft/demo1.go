package main

import (
	"fmt"
	"sync"
)

//1.实现3节点选举
//2.改造代码成分布式选举代码，加入rpc调用
//3. 演示完整代码 自动选主 日志复制

// raft的节点数
const raftCount = 3

//声明leader对象
type Leader struct {
	Term     int //任期
	LeaderId int //leader编号
}

type Raft struct {
	mu              sync.Mutex
	me              int
	currentTerm     int
	votedFor        int
	state           int
	lastMessageTime int64
	currentLeader   int
	message         chan bool
	electCh         chan bool
	heartBeat       chan bool
	heartbeatRe     chan bool
	timeout         int
}

//var (
//	str1,str2 string
//)

func main() {
	str1, str2 := "hello", "world"
	demo(str1, str2)
}

func demo(a ...string) {
	fmt.Println(a)
}
