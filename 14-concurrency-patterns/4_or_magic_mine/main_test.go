package ormagicmine

import (
	"fmt"
	"testing"
	"time"
)

type SIGNAL func(after time.Duration) <-chan struct{}
type HIDDENSIGNAL func() <-chan struct{}

var sig SIGNAL = func(after time.Duration) <-chan struct{} {
	c := make(chan struct{})
	go func() {
		defer close(c)
		<-time.After(after)
	}()
	return c
}

var one HIDDENSIGNAL = func() <-chan struct{} { return sig(1 * time.Second) }
var six HIDDENSIGNAL = func() <-chan struct{} { return sig(6 * time.Second) }
var ten HIDDENSIGNAL = func() <-chan struct{} { return sig(10 * time.Second) }

// type SIGNALS_COLLECTION struct{ signals []HIDDENSIGNAL }

// var testCases []SIGNALS_COLLECTION = []SIGNALS_COLLECTION{
// 	{signals: []HIDDENSIGNAL{one}},
// 	{signals: []HIDDENSIGNAL{one, ten, six, ten}},
// }
//

var TestCases []struct{ function OR } = []struct{ function OR }{
	{function: or},
	{function: orPro},
}

func TestOr0(t *testing.T) {

	for _, testCase := range TestCases {
		f := testCase.function
		start := time.Now()
		<-f(
			1,
			ten(),
			one(),
			one(),
			one(),
			one(),
			one(),
			one(),
			one(),
			one(),
			one(),
			one(),
			one(),
			one(),
			one(),
			one(),
			six(),
			ten(),
		)
		expired := time.Since(start)
		expected, _ := time.ParseDuration("1.2s")

		if expired > expected {
			t.Errorf("Must be ~1 sec., but %v", expired)
		} else {
			fmt.Printf("Well done after %v\n", expired)
		}
	}
}
