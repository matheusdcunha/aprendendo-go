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

func print(c1 <-chan string, c2 <-chan string) {
	for {

		select {
		case msg1 := <-c1:
			log.Println("Canal 1:", msg1)
		case msg2 := <-c2:
			log.Println("Canal 2:", msg2)
		case <-time.After(time.Second * 5):
			log.Println("Tempo esgotado")
			return
		}
	}
}

func main() {

	c1, c2 := make(chan string), make(chan string)

	s1, s2 := time.Duration(6), time.Duration(6)

	go ping(c1, s1)
	go pong(c2, s2)
	print(c1, c2)

}
