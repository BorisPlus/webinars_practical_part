package main

import "fmt"

func main() {
	chanOwner := func() <-chan int {
		results := make(chan int)
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				fmt.Println("Input", i)
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner()
	// close(results)
	consumer(results)
}
