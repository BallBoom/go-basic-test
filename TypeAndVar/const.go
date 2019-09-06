package main

import (
	"fmt"
)

//iota 的值跟顺序有关系，跟多少个没关系
const (
	a = iota
	b
	c = iota
	d
)

//通过iota 求计算机容量单位
const (
	B float32 = 1 << (10 * iota)
	KB
	MB
	GB
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)

}
