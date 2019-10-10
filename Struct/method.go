package main

import (
	"fmt"
)

type A struct {
	Name string
}

type TZ int

func main() {
	a := &A{}
	fmt.Println(&a)
	a.Print()
	fmt.Println(&a)

	var t TZ
	t.Increas(100)
	fmt.Println(t)
}

func (a *A) Print() {
	a.Name = "AA"
	fmt.Println(&a)
}

func (t *TZ) Increas(num int) {
	*t += TZ(num) //TZ的底层虽然是int型的，但是还是有区别，如果要和tz一起运算，就得把传入的参数，强制转换成TZ类型
}
