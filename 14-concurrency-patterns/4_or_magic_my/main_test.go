package main

import (
	"fmt"
	"testing"
	"time"
)

func TestOrFromOne(t *testing.T) {

	sig := func(after time.Duration) <-chan struct{} {
		c := make(chan struct{})
		go func() {
			defer close(c)
			<-time.After(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(1 * time.Second),
	)
	expired := time.Since(start)

	expected, _ := time.ParseDuration("1.2s")
	if expired > expected {
        t.Errorf("Must be ~1 sec., but %v", expired)
    } else {
		fmt.Printf("Well done after %v\n", expired)
	}
}

func TestOrFromTwo(t *testing.T) {

	sig := func(after time.Duration) <-chan struct{} {
		c := make(chan struct{})
		go func() {
			defer close(c)
			<-time.After(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(1*time.Second),
		sig(10*time.Second),
	)
	expired := time.Since(start)

	expected, _ := time.ParseDuration("1.2s")
	if expired > expected {
        t.Errorf("Must be ~1 sec., but %v", expired)
    } else {
		fmt.Printf("Well done after %v\n", expired)
	}
}

func TestOrFromThree(t *testing.T) {

	sig := func(after time.Duration) <-chan struct{} {
		c := make(chan struct{})
		go func() {
			defer close(c)
			<-time.After(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(10*time.Second),
	)
	expired := time.Since(start)

	expected, _ := time.ParseDuration("1.2s")
	if expired > expected {
        t.Errorf("Must be ~1 sec., but %v", expired)
    } else {
		fmt.Printf("Well done after %v\n", expired)
	}
}
	// start = time.Now()
	// <-or(
	// 	sig(5*time.Minute),
	// 	sig(1*time.Second),
	// )
	// fmt.Printf("done after %v\n", time.Since(start))

	// start = time.Now()
	// <-or(
	// 	sig(5*time.Minute),
	// 	sig(1*time.Second),
	// 	sig(10*time.Second),
	// )
	// fmt.Printf("done after %v\n", time.Since(start))

	// start = time.Now()
	// <-or(
	// 	sig(2*time.Hour),
	// 	sig(5*time.Minute),
	// 	sig(1*time.Second),
	// 	sig(1*time.Hour),
	// 	sig(10*time.Second),
	// )
	// fmt.Printf("done after %v\n", time.Since(start))
// }
