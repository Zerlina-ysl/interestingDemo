package channel

import (
	"fmt"
	"testing"
	"time"
)

// channel死锁的几种情况

// 1. 一个通道在一个主线程中同时进行读和写
// 原因：无缓冲chan，发送方和接收方都要准备好才能保证数据的通信。一前一后时，发送方会一直阻塞等待接收方准备就绪
// fatal error: all goroutines are asleep - deadlock!
func TestDeadLock_1(t *testing.T) {
	ch := make(chan int)
	ch <- 100
	name := <-ch
	fmt.Println(name)
}

// 2. 协程开启之前使用通道
// fatal error: all goroutines are asleep - deadlock!
func TestDeadLock_2(t *testing.T) {
	ch := make(chan int)
	// 在协程之前使用chan
	ch <- 100
	go func() {
		num := <-ch
		fmt.Println(num)
	}()

	//ch <- 100 如果在这里使用 就不会死锁
	// sleep是等待协程运行
	time.Sleep(time.Second * 3)
	fmt.Println("finish")
}

func TestDeadLock_3(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for {
			select {
			case num := <-ch1:
				fmt.Println(num)
				ch2 <- 10000
			}
		}
	}()
	for {
		select {
		case num := <-ch2:
			fmt.Println(num)
			ch1 <- 10000
		}
	}

}
func TestDeadLock_4(t *testing.T) {
	ch := make(chan int)
	fmt.Println(<-ch)
}
func TestDeadLock_5(t *testing.T) {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3

}

// panic: send on closed channel
func TestNoDeadLock(t *testing.T) {
	ch := make(chan int, 2)
	close(ch)
	ch <- 1
}
