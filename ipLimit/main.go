package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	mutex sync.Mutex
)

type Ban struct {
	visitIPs map[string]struct{}
}

func NewBan() *Ban {
	return &Ban{visitIPs: map[string]struct{}{}}
}

func main() {
	var success int64 = 0
	ban := NewBan()
	wg := new(sync.WaitGroup)
	// 模拟访问次数
	for i := 0; i < 1000; i++ {
		// 模拟ip
		for j := 0; j < 100; j++ {
			wg.Add(1)
			//ipEnd := j
			go func() {
				defer wg.Done()
				// j不会按序打印
				ip := fmt.Sprintf("192.168.1.%d", j)
				//fmt.Println(j)

				if !ban.visit(ip) {
					//fmt.Println(ip)
					atomic.AddInt64(&success, 1)
				}
			}()
		}
	}
	wg.Wait()
	fmt.Println("success:", success)
}

// visit 判断此时访问的ip是否存在
// 第一次访问会存到map中
func (ban *Ban) visit(ip string) bool {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := ban.visitIPs[ip]; ok {
		return true
	}
	ban.visitIPs[ip] = struct{}{}
	// 在加入ip后异步计算失效时间
	go ban.invalidAfter3Min(ip)
	return false

}

// invalidAfter3Min 开启的异步任务 定时删除
func (ban *Ban) invalidAfter3Min(ip string) {

	time.Sleep(3 * time.Minute)
	mutex.Lock()
	visitedIPs := ban.visitIPs
	delete(visitedIPs, ip)
	ban.visitIPs = visitedIPs
	mutex.Unlock()
}
