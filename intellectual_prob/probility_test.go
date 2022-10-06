package intellectual_prob

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

//不均匀硬币，正面概率是60%，类比一个60%概率返回1的函数，怎样利用这个函数50%的概率返回1？

var coins []int = []int{1, 1, 1, 0, 0}

// 1 正面 0 反面
func foo() int {
	rand.Seed(time.Now().UnixNano())
	// [0,5)
	i := rand.Intn(5)
	return coins[i]

}
func TestUneven(t *testing.T) {
	res := 0
	for i := 0; i < 1000; i++ {
		if foo() == 1 {
			res++
		}
	}
	fmt.Println(res)
}
func Half() int {
	for {
		// 调用foo()两次 出现0,1 和 出现1,0的概率是一样的
		// 对于出现0,0或1,1的情况 舍弃
		a := foo()
		b := foo()
		if a != b {
			return b
		}
	}
}

// 50%
func TestHalf(t *testing.T) {
	res := 0
	for i := 0; i < 10000; i++ {
		if Half() == 1 {
			res++
		}
	}
	fmt.Println(res)
}
