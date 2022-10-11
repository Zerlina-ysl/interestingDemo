package intellectual_prob

import (
	"fmt"
	"math"
	"testing"
)

//给定一个数组，包含从 1 到 N 所有的整数，但其中缺了两个数字。你能在 O(N) 时间内只用 O(1) 的空间找到它们吗？

//https://blog.csdn.net/leelitian3/article/details/109134030

//https://juejin.cn/post/7148461429945794597
var test = []int{1, 2, 5}

//对于数组中的数x1+x2+x3+...Xn==sum-subSum
//1. 将[1,N]取乘积，subMulti为给定数组的乘积，N!=a * b * subMulti
//2. 对[1,N]取平方和，对数组取平方和，他们的差值是a^2+b^2
//3. $(a+b)^2$=$a^2$+$b^2$+2ab
func TestMath(t *testing.T) {
	n := len(test) + 2
	arrayMulti := 1
	arraySum := 0
	nMulti := 1
	nSum := 0
	for i := 0; i < len(test); i++ {

		arrayMulti *= test[i]
		arraySum += test[i] * test[i]
	}
	for i := 1; i <= n; i++ {
		nMulti *= i
		nSum += i * i
	}

	aAddb := int(math.Sqrt(float64(nSum - arraySum + 2*(nMulti/arrayMulti))))
	aMultib := nMulti / arrayMulti
	for i := 1; i < aAddb; i++ {
		j := aAddb - i
		if i*j == aMultib {

			fmt.Println(i, j)
			break
		}
	}

}
