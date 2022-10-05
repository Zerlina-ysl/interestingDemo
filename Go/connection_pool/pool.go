package connection_pool

import (
	"errors"
	"io"
	"sync"
)

type Pool struct {
	m         sync.Mutex     // 使用互斥锁保证release和close的并发安全
	resources chan io.Closer // channel本身并发安全
	closed    bool
}

func (p *Pool) Acquire() (io.Closer, error) {
	r, ok := <-p.resources
	if !ok {
		return nil, errors.New("pool has been closed")
	}
	return r, nil
}
func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		r.Close()
		return
	}
	select {
	case p.resources <- r:
	default:
		r.Close()
	}
}
func (p *Pool) Close() error {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		return nil
	}
	p.closed = true
	close(p.resources)
	for r := range p.resources {
		if err := r.Close(); err != nil {
			return err
		}
	}
	return nil
}
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size too small")
	}
	res := make(chan io.Closer, size)
	for i := 0; i < int(size); i++ {
		c, err := fn()
		if err != nil {
			return nil, err
		}
		res <- c
	}
	return &Pool{
		resources: res,
	}, nil
}
