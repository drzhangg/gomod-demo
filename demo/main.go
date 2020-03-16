package main

import (
	"fmt"
	"gomod-demo/logagent/etcd"
)

func main() {
	fmt.Println(Abc + 1)
	fmt.Println(12)

	etcd.Tstring = "jerry"
	fmt.Println(etcd.Tstring)
	fmt.Println(etcd.Cname)
}
