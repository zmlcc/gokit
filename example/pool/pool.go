package pool

import (
	"fmt"
	"sync"
)

type Value struct {
	string
}

func NewValue(key string) (*Value, error) {
	return &Value{key}, nil
}

var newValue = NewValue

type Pool struct {
	sync.Mutex
	pool    map[string]*Value
	limiter map[string]chan struct{}
}

func NewPool() *Pool {
	return &Pool{
		pool:    make(map[string]*Value),
		limiter: make(map[string]chan struct{}),
	}
}

func (p *Pool) acquire(key string) (*Value, error) {
	// Check to see if there's a pooled connection available. This is up
	// here since it should the the vastly more common case than the rest
	// of the code here.
	p.Lock()
	c := p.pool[key]
	if c != nil {
		p.Unlock()
		return c, nil
	}

	// If not (while we are still locked), set up the throttling structure
	// for this address, which will make everyone else wait until our
	// attempt is done.
	var wait chan struct{}
	var ok bool
	if wait, ok = p.limiter[key]; !ok {
		wait = make(chan struct{})
		p.limiter[key] = wait
	}
	isLeadThread := !ok
	p.Unlock()

	// If we are the lead thread, make the new connection and then wake
	// everybody else up to see if we got it.
	if isLeadThread {
		c, err := newValue(key)
		p.Lock()
		delete(p.limiter, key)
		close(wait)
		if err != nil {
			p.Unlock()
			return nil, err
		}

		p.pool[key] = c
		p.Unlock()
		return c, nil
	}

	// Otherwise, wait for the lead thread to attempt the connection
	// and use what's in the pool at that point.
	<-wait

	// See if the lead thread was able to get us a connection.
	p.Lock()
	if c := p.pool[key]; c != nil {
		p.Unlock()
		return c, nil
	}

	p.Unlock()
	return nil, fmt.Errorf("Lead thread failed to get value")
}
