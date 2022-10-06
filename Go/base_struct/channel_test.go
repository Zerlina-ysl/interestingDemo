package base_struct

import (
	"fmt"
	"sync"
	"testing"
)

func Test_Channel(t *testing.T) {
	// 死循环
	ch := make(chan int)
	go func() {
		ch <- 1
		close(ch)
	}()
	for {
		select {
		case i := <-ch:
			fmt.Println(i)
		default:
			break
		}
	}
}

var (
	a string

	c = make(chan int, 10)
	l sync.Mutex
)

func Test_HappenBefore(t *testing.T) {
	go f()
	<-c
	fmt.Println(a)

}
func f() {
	a = "hello ,world"
	c <- 0
}
func f1() {
	a = "hello world"
	l.Unlock()
}
func Test_Lock(t *testing.T) {
	l.Lock()
	go f1()
	l.Lock()
	fmt.Println(a)
	l.Unlock()
}
