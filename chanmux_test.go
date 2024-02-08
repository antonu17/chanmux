package chanmux

import (
	"testing"
)

func TestChanIntMux(t *testing.T) {

	chans := make([]chan int, 0, 5)
	for i := 0; i < 5; i++ {
		chans = append(chans, make(chan int))
	}
	mux := New[int](chans)

	for i := 0; i < 5; i++ {
		go func(i int) {
			chans[i] <- i
			t.Logf("write to: %d", i)
			close(chans[i])
			t.Logf("close: %d", i)
		}(i)
	}

	for i := range mux.Mux() {
		t.Logf("read from mux: %d", i)
	}

	<-mux.Done()
	t.Logf("mux is done")
}

func TestChanStructMux(t *testing.T) {
	chans := make([]chan struct{}, 0, 5)
	for i := 0; i < 5; i++ {
		chans = append(chans, make(chan struct{}))
	}
	mux := New[struct{}](chans)

	for i := 0; i < 5; i++ {
		go func(i int) {
			close(chans[i])
			t.Logf("close: %d", i)
		}(i)
	}

	<-mux.Done()
	t.Logf("mux is done")
}
