package main

import (
	"fmt"
)

func main() {
	//匿名函数
	a := func() {
		fmt.Println("func a")
	}
	a()

	f := closure(5)
	fmt.Println(f(3))
	fmt.Println(f(4))

}

//函数闭包：闭包函数的参数，在代码块里调用是直接引用，而不是值拷贝
func closure(x int) func(int) int {
	fmt.Printf("%p\n", &x) //两次打印出来的额地址都是一样的
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}

}
