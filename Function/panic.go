package main

import (
	"fmt"
)

func main() {
	A()
	B()
	C()
}

func A() {
	fmt.Println("Func A")
}

//引发panic，程序出现panic时就会停止运行，因此C()不会运行
//func B() {
//panic("Panic in B")
//}

//使用recover 修复错误，出现recover时函数从panic状态修复，还会继续调用
func B() {
	defer func() { //defer 要再panic之前注册，在发生panic之后发现有defer，执行defer
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()
	panic("Panic in B")
}

func C() {
	fmt.Println("Func C")
}
