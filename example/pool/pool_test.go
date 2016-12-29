package pool

import (
	"sync"
	"testing"
	"time"
)

func Test_acquire(t *testing.T) {
	old := newValue
	defer func() {
		newValue = old
	}()

	callCount := 0
	newValue = func(key string) (*Value, error) {
		callCount += 1
		time.Sleep(3 * time.Second)
		t.Log("Call Count:", callCount)
		return old(key)
	}

	p := NewPool()

	callNumber := 5
	wg := new(sync.WaitGroup)
	wg.Add(callNumber)
	for _ = range make([]struct{}, callNumber) {
		go func() {
			t.Log(p.acquire("test"))
			wg.Done()
		}()
	}
	wg.Wait()

}
