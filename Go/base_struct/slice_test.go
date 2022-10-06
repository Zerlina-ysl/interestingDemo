package base_struct

import (
	"fmt"
	"testing"
)

func Test_Slice(t *testing.T) {
	var s []int
	for i := 0; i < 3; i++ {
		s = append(s, i)
	}
	//modifySlice1(s)
	//modifySlice2(s)
	//modifySlice3(s)
	modifySlice4(s)
	fmt.Println(s)
}
func modifySlice1(s []int) {
	// [1024 1 2]
	s[0] = 1024
}
func modifySlice2(s []int) {
	s = append(s, 2048)
	// [1024 1 2]
	// 传入的s和方法中的s是不同的对象
	// type base_struct struct {
	//	array unsafe.Pointer
	//	len   int	总容量
	//	cap   int 	当前长度
	//}
	// 不同的slice struct 但是指针指向同一块区域
	s[0] = 1024
}
func modifySlice3(s []int) {
	// [0 1 2]
	// 	初始分配slice len=4 cap=3
	s = append(s, 2048)
	// 此时slice需要扩容 重新分配内存 因此不和之前的slice指向相同的地址空间
	s = append(s, 4096)
	// 方法内修改，方法外无法观察
	s[0] = 1024
}
func modifySlice4(s []int) {
	// [1024 1 2]
	s[0] = 1024
	s = append(s, 2048)
	s = append(s, 4096)
}
