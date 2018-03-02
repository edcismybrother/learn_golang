package main 

import (
	"time"
	"math/rand"
	"sync"
	"log"
)

/*
Channel关闭原则
不要在消费端关闭channel，不要在有多个并行的生产者时对channel执行关闭操作。
*/

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	
	// ...
	const MaxRandomNumber = 100000
	const NumReceivers = 100
	
	wgReceivers := sync.WaitGroup{}
	wgReceivers.Add(NumReceivers)
	
	// ...
	dataCh := make(chan int, 100)
	
	// the sender
	go func() {
		for {
			if value := rand.Intn(MaxRandomNumber); value == 0 {
				// The only sender can close the channel safely.
				close(dataCh)
				return
			} else {			
				dataCh <- value
			}
		}
	}()
	
	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wgReceivers.Done()
			
			// Receive values until dataCh is closed and
			// the value buffer queue of dataCh is empty.
			for value := range dataCh {
				log.Println(value)
			}
		}()
	}
	
	wgReceivers.Wait()
}

