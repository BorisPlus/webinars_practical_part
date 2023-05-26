package ormagicmine

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
