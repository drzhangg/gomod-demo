package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
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
	mu              sync.Mutex //锁
	me              int        //节点编号
	currentTerm     int        //当前任期
	votedFor        int        //为哪个节点投票
	state           int        //3个状态 0 follower 1 candidate 2 leader
	lastMessageTime int64      // 发送最后一条数据的时间
	currentLeader   int        // 设置当前节点的领导
	message         chan bool  // 节点间发消息的通道
	electCh         chan bool  // 选举通道
	heartBeat       chan bool  // 心跳信号的通道
	heartbeatRe     chan bool  // 返回心跳信号的通道
	timeout         int        // 超时时间 15~30 ms
}

// 0 还没上任， -1 没有编号
var leader = Leader{
	Term:     0,
	LeaderId: -1,
}

func main() {
}

// me入参是设置当前节点编号
func Make(me int) *Raft {
	rf := &Raft{}
	rf.me = me
	// -1代表谁都不投，此时节点刚创建
	rf.votedFor = -1
	// 0 follower
	rf.state = 0
	rf.timeout = 0
	rf.currentLeader = -1
	//节点任期
	rf.setTerm(0)

	rf.message = make(chan bool)
	rf.electCh = make(chan bool)     // 选举通道
	rf.heartBeat = make(chan bool)   // 心跳信号的通道
	rf.heartbeatRe = make(chan bool) // 返回心跳信号的通道

	//选举的协程

	//心跳检测的协程

}

//实现选主的逻辑
func (rf *Raft) election_one_round(leader *Leader) bool {
	//定义超时
	var timeout int64
	timeout = 100
	//投票数量
	var vote int
	//定义是否开始心跳信号的产生
	var triggerHeartbeat bool
	//时间
	last := millisecond()

	success := false

	//给当前节点变成candidate
	rf.mu.Lock()
	rf.becomeCandidate()
	rf.mu.Unlock()
	fmt.Println("start election leader")

	for {
		//遍历所有节点拉选票
		for i := 0; i < raftCount; i++ {
			if i != rf.me {
				//拉选票
				go func() {
					if leader.LeaderId < 0 {
						//设置投票
						rf.electCh <- true
					}
				}()
			}
		}
		//设置投票数量
		vote = 1
		//遍历
		for i := 0; i < raftCount; i++ {
			//计算投票数量
			select {
			case ok := <-rf.electCh:
				if ok {
					//投票数量+1
					vote++
					//若选票个数，大于节点个数/2，则成功
					success = vote > raftCount/2
					if success && !triggerHeartbeat {
						//变化成主节点，选主成功了
						//开始触发心跳信号检测
						triggerHeartbeat = true

						rf.mu.Lock()
						//变主
						rf.becomeLeader()
						rf.mu.Unlock()

						//由leader向其他节点发送心跳信号
						rf.heartBeat <- true
						fmt.Println(rf.me, "号节点成为了leader")
						fmt.Println("leader开始发送心跳信号了")
					}
				}
			}
		}
	}
}

func (rf *Raft) setTerm(term int) {
	rf.currentTerm = term
}

// 随机值
func randRange(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

// 获取当前时间，发送最后一条数据的时间
func millisecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// 修改状态candidate
func (rf *Raft) becomeCandidate() {
	rf.state = 1 //1为candidate状态
	// 将当前节点任期加1
	rf.setTerm(rf.currentTerm + 1)
	rf.votedFor = rf.me //为当前节点投票
	rf.currentLeader = -1
}

// 修改状态leader
func (rf *Raft) becomeLeader() {
	rf.state = 2 //2为leader的状态
	rf.currentLeader = rf.me
}
