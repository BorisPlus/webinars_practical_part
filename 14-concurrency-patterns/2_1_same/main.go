package main

import (
	"fmt"
	"sync"
)

func main() {

	data := []int{1, 2, 3, 4, 5}

	chanOwner := func(srcData []int) <-chan int {
		results := make(chan int)
		var wg sync.WaitGroup
		wg.Add(5)
		for _, k := range data {
			go func(k int) {
				defer wg.Done()
				for i := 0; i < 5; i++ {
					fmt.Printf("Input: %d\n", i + k * 10)
					results <- (i + k * 10)
				}
			}(k)
		}
		go func() {
		wg.Wait()
		defer close(results)
		}()
		return results
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner(data)
	consumer(results)
}
