package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	var url = "https://movie.douban.com/top250"

	resp, _ := http.Get("https://movie.douban.com/top250")
	fmt.Println(resp)

	buf := make([]byte, 1024)
	var re string
	for {
		n, err := resp.Body.Read(buf)
		re += string(buf[:n])
		if err != nil {
			if err == io.EOF {
				break
			} else {
				break
			}
		}
	}
	fmt.Println(re)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(request)

	//fmt.Println(request.Response.Body)
	//
	//body, _ := ioutil.ReadAll(request.Response.Body)
	//fmt.Println(string(body))
}
