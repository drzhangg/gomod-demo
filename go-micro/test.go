package main

import (
	"github.com/micro/go-micro/registry"
	"log"
)

func main() {

	services, err := registry.GetService("prodservice1")
	if err != nil {
		log.Fatal(err)
	}

}
