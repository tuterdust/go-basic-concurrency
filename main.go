package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println("Hello, Test Go Channel")

	rand.Seed(time.Now().UnixNano())
	numOrder := 10
	ch := make(chan int, 10)
	var totalTime int
	wg := &sync.WaitGroup{}

	wg.Add(numOrder)
	for i := 0; i < numOrder; i++ {
		go buyItem(wg, &ch, i+1)
	}

	wg.Wait()
	close(ch)
	for e := range ch {
		totalTime += e
	}

	fmt.Printf("Finished Execution with a total time of %d seconds.", totalTime)
}

func buyItem(wg *sync.WaitGroup, ch *chan int, orderNum int) {
	sleepTime := rand.Intn(5)
	time.Sleep(time.Duration(sleepTime) * time.Second)
	fmt.Printf("Order %d is completed.\n", orderNum)
	*ch <- sleepTime
	wg.Done()
}

