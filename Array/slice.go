package main

import (
	"fmt"
)

func main() {

	var s1 []int
	fmt.Println(s1)

	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := a[5:10] // a[5,6,7,8,9] 只包含第一个索引值，后面索引值不包含
	//s2 := a[5:] 从下标5开始取;  a[:5]  取到下标4
	fmt.Println(s2)

	s3 := make([]int, 3, 5) //使用make创建slice有类型零值
	fmt.Printf("%v %p\n", s3, s3)
	s3 = append(s3, 5, 6, 7, 8) //slice 使用append时，如果超过了一开始创建的容量，则会重新分配一个数组并指向新的slice
	fmt.Printf("%v %p", s3, s3)

}
