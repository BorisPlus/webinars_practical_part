package main

// https://github.com/OtusGolang/webinars_practical_part/blob/master/14-concurrency-patterns/4_or_magic/main.go
// go test -v -count=1 -race -timeout=1m ./

func or(channels ...<-chan struct{}) <-chan struct{} {
	if len(channels) == 0 {
		return nil
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
