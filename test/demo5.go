package main

func main() {
	ch := make(chan int)
	go put(ch)
	go get(ch)
}

func put(ch chan int) {
	ch <- 1
}

func get(ch chan int)  {
	<- ch
}

