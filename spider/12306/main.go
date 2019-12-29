package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//var url = "https://kyfw.12306.cn/otn/leftTicket/init?linktypeid=dc&fs=%E5%B9%BF%E5%B7%9E,GZQ&ts=%E5%91%A8%E5%8F%A3,ZKN&date=2019-12-29&flag=N,N,Y"
	//var url = "https://kyfw.12306.cn/"
	var url = "https://kyfw.12306.cn/otn/resources/js/framework/station_name.js?station_version=1.8994"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	datas, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	_ = ioutil.WriteFile("12306.txt",datas,0666)
	
	fmt.Println(string(datas))
}
