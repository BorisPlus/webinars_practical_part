package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(done <-chan struct{}, strings <-chan string) <-chan int {
		terminated := make(chan int)
		go func() {
			defer func() {
				fmt.Println("doWork exited.")
				close(terminated)
			}()
			for {
				select {
				case s := <-strings:
					fmt.Println("Println", s)
				case <-done:
					return 
				}
			}
		}()
		return terminated
	}

	done := make(chan struct{})
	terminated := doWork(done, nil)

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()

	<-terminated
	fmt.Println(<-terminated)
	fmt.Println("Done.")
}
