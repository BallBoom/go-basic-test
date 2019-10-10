package main

import (
	"fmt"
)

func main() {
	//匿名函数
	/*	a := func() {
			fmt.Println("func a")
		}
		a()

		f := closure(5)
		fmt.Println(f(3))
		fmt.Println(f(4))*/

	bb()

}

//函数闭包：闭包函数的参数，在代码块里调用是直接引用，而不是值拷贝
func closure(x int) func(int) int {
	fmt.Printf("%p\n", &x) //两次打印出来的额地址都是一样的
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}

}

func bb() {
	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		//这边的i是参数，是值拷贝类型，跟for循环中的i是同值
		defer fmt.Println("defer i =", i)

		//同下道理，如果是有参数，那么参数可以取得值拷贝，即与外层的i一致
		defer func(i int) { fmt.Println("defer_closure i=", i) }(i)
		//闭包：如果匿名函数没有定义参数，那么闭包函数里面的参数是从外层取来的，得到的只是“引用地址”，既是for循环中的i
		fs[i] = func() { fmt.Println("closure i = ", i) }
	}

	for _, f := range fs {
		f()
	}

}
