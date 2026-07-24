package main

import (
	"fmt"
	"sync"
	"time"
)

type Dispatcher struct {
	dataChan chan int
	wg       sync.WaitGroup
}

func (d *Dispatcher) AddConsumer(id int) {
	d.wg.Add(1)
	go func() {
		defer d.wg.Done()
		for val := range d.dataChan {
			fmt.Printf("Consumer %d processed: %d\n", id, val)
			time.Sleep(time.Millisecond * 200)
		}
		fmt.Printf("Consumer %d shutting down\n", id)
	}()
}

func main() {
	d := Dispatcher{dataChan: make(chan int)}

	// Start with 2 consumers
	d.AddConsumer(1)
	d.AddConsumer(2)

	// Producer
	go func() {
		for i := 1; i <= 10; i++ {
			d.dataChan <- i
			if i == 5 {
				fmt.Println("--- Adding Consumer 3 Dynamically ---")
				d.AddConsumer(3)
			}
		}
		close(d.dataChan)
	}()

	d.wg.Wait()
}
