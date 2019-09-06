package main

import (
	"fmt"
	"strconv"
)

//类型别名
type (
	s  string
	it int
)

//全局变量声明
/*var (
	a int16
	b float32 = 20.3
)*/

func main() {

	var a1 int = 65
	var b1 string
	b1 = string(a1)
	fmt.Println("b1=", b1) //这边b1= 'A' 原因是string()函数会默认把数字转换成ASC码
	b1 = strconv.Itoa(a1)
	fmt.Println("b1=", b1) //这边b1= '65' strconv()函数适用于int类型转换成字符串
	a1, _ = strconv.Atoi(b1)
	fmt.Println("a1=", a1)
	//fmt.Println(math.MinInt8)
	//var it t

	//test()

}
