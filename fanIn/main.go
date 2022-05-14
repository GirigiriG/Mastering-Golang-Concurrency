package main

import (
	"fmt"
)

func main() {
	fanCh := fanIn(fanInGenerator(23), fanInGenerator(66))
	for cc := range fanCh {
		fmt.Println(cc)
	}
}

func fanInGenerator(data int32) <-chan int32 {
	out := make(chan int32)
	go func(data int32) {
		for i := 0; i < 5; i++ {
			out <- data
		}
		close(out)
	}(data)

	return out
}

func fanIn(gens ...<-chan int32) <-chan int32 {
	fmt.Println(len(gens))
	out := make(chan int32)

	for _, ch := range gens {
		go func(cc <-chan int32) {
			for data := range cc {
				out <- data
			}
			close(out)
		}(ch)
	}

	return out
}
