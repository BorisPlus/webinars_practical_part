package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan struct{}) <-chan struct{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0] // <-or(ch) == <-ch
	}

	orDone := make(chan struct{})
	go func() {
		defer close(orDone)
		select {
		case <-channels[0]:
		case <-or(append(channels[1:], orDone)...):
		}
	}()
	return orDone
}

func main() {
	sig := func(after time.Duration) <-chan struct{} {
		c := make(chan struct{})
		go func() {
			defer close(c)
			<-time.After(after)
		}()
		return c
	}

	start := time.Now()
	<-or(make(<-chan struct{}))
	fmt.Printf("done after %v\n", time.Since(start))

	start = time.Now()
	<-or(
		sig(1*time.Second),
	)
	fmt.Printf("done after %v\n", time.Since(start))

	start = time.Now()
	<-or(
		sig(5*time.Minute),
		sig(1*time.Second),
	)
	fmt.Printf("done after %v\n", time.Since(start))

	start = time.Now()
	<-or(
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(10*time.Second),
	)
	fmt.Printf("done after %v\n", time.Since(start))

	start = time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(10*time.Second),
	)
	fmt.Printf("done after %v\n", time.Since(start))
}
