package main

import (
	"fmt"
)

func main() {
	//定义和初始化
	a := [...]int{12: 1}
	a[3] = 5
	b := [...]int{1, 2, 8: 11}
	//定义指针 取a的地址
	var p *[13]int
	p = &a
	//使用new创建数组
	pr := new([5]int)
	pr[3] = 5
	//定义多维数组 在一维数组中使用的规则 在多维数组中也可以使用
	arr := [2][3]int{
		{1, 2, 3},
		{3, 2, 1},
	}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(p)  //这里P的值和a一致，多了取地址符号 ：&[0 0 0 0 0 0 0 0 0 0 0 0 1]
	fmt.Println(pr) //&[0 0 0 0 0]
	fmt.Println(arr)
	maopao()
}

func maopao() {
	// 冒泡排序
	a := [...]int{2, 8, 5, 9, 3}
	fmt.Println(a)

	num := len(a)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[j] > a[i] {
				a[j], a[i] = a[i], a[j]
			}
		}
	}
	fmt.Println(a)
}
