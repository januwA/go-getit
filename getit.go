package getit

import (
	"fmt"
	"sync"
)

type Getit[T any] struct {
	lazy bool

	cache map[string]any

	cacheMutex sync.RWMutex

	joinCallback func(serv T) T
}

func (my *Getit[T]) New(lazy bool) *Getit[T] {
	my.lazy = lazy
	my.cache = map[string]any{}
	return my
}

func (my *Getit[T]) SetJoinCallback(cb func(serv T) T) {
	my.joinCallback = cb
}

func (my *Getit[T]) callJoinCallback(serv T) T {
	if my.joinCallback == nil {
		return serv
	} else {
		return my.joinCallback(serv)
	}
}

// Register in order
func (my *Getit[T]) Register(s ...T) *Getit[T] {
	for _, serv := range s {
		k := fmt.Sprintf("%T", serv)
		my.cache[k] = my.callJoinCallback(serv)
	}
	return my
}

func (my *Getit[T]) Get(serv_ptr any) any {
	if my.lazy {
		my.cacheMutex.Lock()
		defer my.cacheMutex.Unlock()
	}

	k := fmt.Sprintf("%T", serv_ptr)
	s := my.cache[k]

	if s != nil {
		return s
	}

	if my.lazy {
		my.Register(serv_ptr.(T))
		return my.cache[k]
	} else {
		panic(fmt.Sprintf("%s unregistered!", k))
	}
}
