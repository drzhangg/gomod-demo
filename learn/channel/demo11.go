package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var urls = []string{"http://www.baidu.com", "http://www.bilibili.com"}

	var wg sync.WaitGroup

	for _, v := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, _ := http.Get(v)
			fmt.Println(resp.Header)
		}(v)
	}
	wg.Wait()
}
