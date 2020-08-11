package main

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	"log"
	"strconv"
)

type Person struct {
	Id       int
	Name     string
	Age      int
	NickName string
}

func main() {
	config := nsq.NewConfig()
	w, err := nsq.NewProducer("47.103.9.218:4150", config)
	if err != nil {
		log.Panic("Could not create producer.")
	}
	defer w.Stop()
	for i := 0; i < 100; i++ {
		p := &Person{}
		p.Id = i
		p.Name = "Jack" + strconv.Itoa(i)
		p.NickName = "Luo" + strconv.Itoa(i)
		p.Age = i
		info, jerr := json.Marshal(p)
		err := w.Publish("write_test", info)
		if err != nil || jerr != nil {
			log.Panic("Could not connect.")
		}
	}
	w.Stop()
}
