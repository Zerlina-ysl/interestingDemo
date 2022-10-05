package sync_mutex

import (
	"sync"
	"testing"
	"time"
)

var lockA = &sync.Mutex{}
var lockB = &sync.Mutex{}

// A+B
func A() {
	lockA.Lock()
	defer lockA.Unlock()
	time.Sleep(1 * time.Second)
	lockB.Lock()
	defer lockB.Unlock()
}

// B+A
func B() {
	lockB.Lock()
	defer lockB.Unlock()
	time.Sleep(1 * time.Second)
	lockA.Lock()
	defer lockA.Unlock()
}

func TestAPB(t *testing.T) {
	//!  all goroutines are asleep - deadlock!
	go A()
	//B()
	go B()
}
