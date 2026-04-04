package main

import (
	"log"
	"time"
)

func ping(c chan<- string, t time.Duration) {
	for {
		time.Sleep(time.Second * t)
		c <- "ping"
	}
}

func pong(c chan<- string, t time.Duration) {
	for {
		time.Sleep(time.Second * t)
		c <- "pong"
	}
}

func print(c <-chan string) {
	for {
		msg := <-c
		log.Println(msg)
	}
}

func gerarNumeros(ch chan<- int) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second * 1)
		ch <- i
	}
	close(ch)
}

func main() {

	// c := make(chan string)
	c2 := make(chan int)
	// seconds := time.Duration(1)

	// go ping(c, seconds)
	// go pong(c, seconds)
	// go print(c)

	go gerarNumeros(c2)
	for valor := range c2 {
		log.Println(valor)
	}

	// var ok string
	// fmt.Scanf("%s", &ok)
}
