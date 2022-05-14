package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan interface{})
	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		for i := 0; i < 100; i++ {
			ch <- i
		}
		defer wg.Done()
		close(ch)
	}()
	wg.Wait()

	for data := range ch {
		fmt.Println(data)
	}
}
