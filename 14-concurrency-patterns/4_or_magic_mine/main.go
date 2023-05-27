package ormagicmine

import "fmt"

type OR func(deep int, channels ...<-chan struct{}) <-chan struct{}

func or(deep int, channels ...<-chan struct{}) <-chan struct{} {
	fmt.Println(len(channels), deep)
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0] // <-or(ch) == <-ch
	}

	orDone := make(chan struct{})
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(deep+1, append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}

func or2(deep int, channels ...<-chan struct{}) <-chan struct{} {
	fmt.Println(len(channels), deep)
	if len(channels) == 0 {
		return nil
	}
	orDone := make(chan struct{})
	go func() {
		defer close(orDone)
		select {
		case <-channels[0]:
		case <-or2(deep+1, append(channels[1:], orDone)...):
		}
	}()
	return orDone
}

func orPro(deep int, channels ...<-chan struct{}) <-chan struct{} {
	fmt.Println(len(channels), deep)
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0] // <-or(ch) == <-ch
	case 2:
		select {
		case <-channels[0]:
		case <-channels[1]:
		} 
	case 3:
		select {
		case <-channels[0]:
		case <-channels[1]:
		case <-channels[2]:
		} 
	}

	orDone := make(chan struct{})
	go func() {
		defer close(orDone)
		select {
		case <-channels[0]:
		case <-channels[1]:
		case <-channels[2]:
		case <-channels[3]:
		case <-or(deep+1, append(channels[4:], orDone)...):
		}
	}()
	return orDone
}
