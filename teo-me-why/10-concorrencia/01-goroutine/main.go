package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func diz(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	for range 4 {
		log.Println(s)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var wg sync.WaitGroup

	for i := range 3 {
		wg.Add(1)
		go diz(fmt.Sprintf("%d", i), &wg)
	}

	wg.Wait()
	log.Println("Fim do programa!!")
}
