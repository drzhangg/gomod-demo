package main

import (
	"gomod-demo/learn/pprof/data"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		for {
			log.Println(data.Add("https://github.com/drzhangg"))
		}
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}
