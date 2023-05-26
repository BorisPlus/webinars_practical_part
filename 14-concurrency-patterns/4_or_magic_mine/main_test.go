package ormagicmine

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

//
// === RUN   TestOrFromOne
// Well done after 1.000203168s
// --- PASS: TestOrFromOne (1.00s)
// === RUN   TestOrFromTwo
// Well done after 1.000208127s
// --- PASS: TestOrFromTwo (1.00s)
// === RUN   TestOrFromThree
// Well done after 1.000244864s
// --- PASS: TestOrFromThree (1.00s)
// PASS
// ok      github.com/OtusGolang/webinars_practical_part/14-concurrency-patterns/4_or_magic_mine   3.228s
