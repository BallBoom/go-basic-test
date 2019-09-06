package main

import (
	"fmt"
)

func main() {

	a := 1
	switch {
	case a > 0:
		fmt.Println("a=0")
		fallthrough
	case a >= 1:
		fmt.Println("a=1")
	case a > 2:
		fmt.Println("a=2")
	}
	//	test()
	test2()

}

func test() {
	//LABLE1:
	for {
		for i := 0; i < 10; i++ {
			if i > 2 {
				//break LABLE1  break 是跳出
				fmt.Println("lab")
				goto LABLE1 //goto 是调整位置

			}
		}
	}
LABLE1:
	fmt.Println("aaaaa")
}

func test2() {
LABLE2:
	for i := 0; i < 10; i++ {
		for {
			continue LABLE2 //跳过后面的语句 继续循环
			fmt.Println("ok")
			//goto LABLE2    直接变成死循环 一直执行LABLE2
		}
	}
}
