package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	var url = "https://kyfw.12306.cn/otn/leftTicket/init?linktypeid=dc&fs=%E5%B9%BF%E5%B7%9E,GZQ&ts=%E5%91%A8%E5%8F%A3,ZKN&date=2020-01-20&flag=N,N,Y"
	//var url = "https://kyfw.12306.cn/otn/leftTicket/init?linktypeid=dc&fs=%E5%B9%BF%E5%B7%9E,GZQ&ts=%E5%91%A8%E5%8F%A3,ZKN&date=2020-01-01&flag=N,N,Y"
	//var url = "https://kyfw.12306.cn/otn/resources/js/framework/station_name.js?station_version=1.8994"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	request.Header.Set("Accept","*/*")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")

	//resp, err := http.Get(url)
	//if err != nil {
	//	panic(err)
	//}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Header)
	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)

	datas, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	_ = ioutil.WriteFile("12306.txt", datas, 0666)

	fmt.Println(string(datas))
}


//func getUrlHtml(strUrl string, postDict map[string]string) (bool, string) {
//	var respHtml string = ""
//	CurCookies := nil
//	httpClient := &http.Client{
//		Jar: gCurCookieJar,
//	}
//	var httpReq *http.Request
//	if nil == postDict {
//		httpReq, _ = http.NewRequest("GET", strUrl, nil)
//	} else {
//		postValues := url.Values{}
//		for postKey, PostValue := range postDict {
//			postValues.Set(postKey, PostValue)
//		}
//		postDataStr := postValues.Encode()
//		postDataBytes := []byte(postDataStr)
//		postBytesReader := bytes.NewReader(postDataBytes)
//		httpReq, _ = http.NewRequest("POST", strUrl, postBytesReader)
//		httpReq.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36")
//		httpReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
//		for _, c := range gCurCookies {
//			httpReq.AddCookie(c)
//		}
//	}
//	httpResp, err := httpClient.Do(httpReq)
//	if err != nil {
//
//		fmt.Print("Error 请检查网路\n  http get strUrl=%s response error=%s\n", strUrl, err.Error())
//		return false, ""
//	}
//	defer httpResp.Body.Close()
//	body, errReadAll := ioutil.ReadAll(httpResp.Body)
//	if errReadAll != nil {
//		fmt.Print("Error  :  get response for strUrl=%s got error=%s\n", strUrl, errReadAll.Error())
//		return false, ""
//	}
//	gCurCookies = gCurCookieJar.Cookies(httpReq.URL)
//	respHtml = string(body)
//	return true, respHtml
//}
