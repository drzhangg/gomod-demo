package registry

//抽象服务
type Service struct {
	Name  string `json:"name"`  //服务名
	Nodes []Node `json:"nodes"` //节点列表
}

//单个服务节点的抽象
type Node struct {
	Id     string `json:"id"`
	Ip     string `json:"ip"`
	Port   int    `json:"port"`
	Weight int    `json:"weight"`
}
