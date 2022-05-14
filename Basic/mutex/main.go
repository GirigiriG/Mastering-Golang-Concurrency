package main

import "fmt"

func main() {

	d := fileReader("gideon")

	for data := range d {
		fmt.Println(data)
	}

}

func fileReader(data string) <-chan string {
	ch := make(chan string)
	go func() {

		for i := 0; i < 10; i++ {
			ch <- fmt.Sprintf("lines %d: %v", i, data)
		}
		close(ch)
	}()

	return ch
}
