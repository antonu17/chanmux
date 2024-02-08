package chanmux

import (
	"sync"
)

type ChanMuxer[T interface{}] interface {
	Mux() <-chan T
	Done() <-chan T
}

type chanMux[T interface{}] struct {
	mux  chan T
	done chan T
}

func New[T interface{}](cs []chan T) ChanMuxer[T] {
	mux := make(chan T)
	done := make(chan T)

	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(c <-chan T) {
			for v := range c {
				mux <- v
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(mux)
		close(done)
	}()

	return &chanMux[T]{
		mux:  mux,
		done: done,
	}
}

func (cs *chanMux[T]) Mux() <-chan T {
	return cs.mux
}

func (cs *chanMux[T]) Done() <-chan T {
	return cs.done
}
