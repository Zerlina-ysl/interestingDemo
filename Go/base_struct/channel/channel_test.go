package channel

import (
	"fmt"
	"sync"
	"testing"
)

// 使⽤两个goroutine交替打印序列，⼀个goroutine打印数字， 另外⼀个goroutine打印字⺟， 最终效果如下：

// 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728

// https://juejin.cn/post/6965840516533452808
func TestLoopPrint(t *testing.T) {
	// 使用两个无缓冲的channel控制协程
	number, letter := make(chan bool), make(chan bool)

	var wg sync.WaitGroup
	// 等待number的管道通知->打印数字->通知number管道
	go func() {
		i := 1
		for {
			<-number
			fmt.Printf("%d%d", i, i+1)
			i += 2
			letter <- true
		}
	}()
	wg.Add(1)
	go func() {
		str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			<-letter
			if i > len(str)-2 || i < 0 {
				wg.Done()
				return
			}
			fmt.Print(str[i : i+2])
			i += 2
			number <- true
		}
	}()
	number <- true
	// 阻塞直到计数器为0
	wg.Wait()
}
