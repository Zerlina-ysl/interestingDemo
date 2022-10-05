package timeout_control

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/*------利用time.After实现超时控制---------*/
func Test_Timer(t *testing.T) {
	fmt.Println(time.Now())
	// func After(d Duration) <-chan Time {
	//	return NewTimer(d).C
	// }
	x := <-time.After(3 * time.Second)
	fmt.Println(x)
}
func Test_After(t *testing.T) {
	ch := make(chan struct{}, 1)
	go func() {
		fmt.Println("do")
		time.Sleep(4 * time.Second)
		ch <- struct{}{}
	}()
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("time out")
	case <-ch:
		fmt.Println("done")
	}
}

/*-----利用context实现超时控制-----*/
func Test_Context(t *testing.T) {
	ch := make(chan string)
	// 返回一个具有超时功能的context
	timeout, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	go func() {
		time.Sleep(time.Second * 4)
		ch <- "done"
	}()
	select {
	case res := <-ch:
		fmt.Println(res)
		// Done() <-chan struct{}
	case <-timeout.Done():
		// timeout.Err可以得到channel关闭的原因
		fmt.Println("timeout", timeout.Err())
	}
}

/*--------利用time.sleep实现超时控制--------*/
func Test_TimeSleep(t *testing.T) {
	timeout := make(chan bool, 1)
	ch := make(chan int, 1)
	go func() {
		time.Sleep(3 * time.Second)
		// 3s后向timeout的channel中发送值
		timeout <- true
	}()
	select {
	// 此时拿到了值，说明已超时
	case <-timeout:
		fmt.Println("timeout")
	case <-ch:
		fmt.Println("no timeout")
	}
}

func Test_NewTimer(t *testing.T) {
	ch := make(chan bool, 1)
	defer close(ch)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- true
	}()
	timer := time.NewTimer(1 * time.Second)
	defer timer.Stop()
	select {
	case <-ch:
		fmt.Println("read ch")
	case <-timer.C:
		fmt.Println("time out")
	}
}
