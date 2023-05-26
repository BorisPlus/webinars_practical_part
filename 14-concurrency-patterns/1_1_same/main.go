package main

import "fmt"

func main() {

	data := []int{1, 2, 3, 4}

	handleData := make(chan int)

	go func(handleDataIn chan<- int) {
		defer close(handleData)
		for i := range data {
			handleDataIn <- data[i]
		}
	}(handleData)

	for num := range handleData {
		fmt.Println(num)
	}
}