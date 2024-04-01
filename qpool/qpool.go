package qpool

import "sync"

type QPool[T any] struct {
	sync.Pool
}

func NewQPool[T any](newFunc func() T) *QPool[T] {
	return &QPool[T]{
		Pool: sync.Pool{New: func() any {
			return newFunc()
		}},
	}
}

func (p *QPool[T]) Get() T {
	return p.Pool.Get().(T)
}

func (p *QPool[T]) Put(t T) {
	p.Pool.Put(t)
	return
}
