package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func main() {
	fasthttp.ListenAndServe(":8087", Upload)
}

func Upload(ctx *fasthttp.RequestCtx) {
	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(file.Filename)
}
