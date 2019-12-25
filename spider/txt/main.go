package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
)

func main() {
	var url = "http://www.siluke.cc/html/43/37539.html"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("resp get err:", err)
	}

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	document.Find("#content").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())

		var d1 = []byte(s.Text())
		err := ioutil.WriteFile("./output2.txt", d1, 0666)
		if err != nil {
			fmt.Println(err)
		}
	})

	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))

}
