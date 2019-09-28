package main

import (
	"fmt"
	"sort"
)

func main() {
	//定义一个map，多种定义方法
	var m map[int]int = make(map[int]int, 5)
	//m = make(map[int]int, 5)
	// m := make(map[int]int, 5)  简单map
	m[1] = 100
	m[2] = 200
	fmt.Println(m)
	//删除元素
	delete(m, 1)
	fmt.Println(m)

	//定义复杂map
	var b map[int]map[int]string
	b = make(map[int]map[int]string)
	a, ok := b[1][1] //使用多返回值 可判断是否初始化
	if ok == false {
		b[1] = make(map[int]string)
	}
	b[1][1] = "OK"
	a, ok = b[1][1]
	fmt.Println(a, ok)

	//创建slice类型map
	sm := make([]map[int]string, 3)
	//使用range 对其赋值，如果不使用slice的下标赋值 则 只是值拷贝，并不会对原slice的值有改变
	//如果是map的话 需要用key赋值
	for _, v := range sm {
		v = make(map[int]string, 1)
		v[1] = "ok"
		fmt.Println(v) //map[1:ok]
	}
	fmt.Println(sm) //迭代完之后 sm里的map还是空的,所以使用下标对原slice赋值

	for i := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "ok"
		fmt.Println(sm[i])
	}
	fmt.Println(sm)

	//间接排序map的key
	map1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	slice1 := make([]int, len(map1))
	//map里没有索引 所以自建索引
	i := 0
	for k, _ := range map1 { //我发现 如果不迭代value的话 每次key都是按顺序的
		slice1[i] = k
		i++
	}
	// 对slice进行排序
	sort.Ints(slice1)
	fmt.Println(slice1)

	//key value 转换
	map2 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	map3 := make(map[string]int, len(map2))

	fmt.Println("转换前：", map2)
	for k, v := range map2 {
		map3[v] = k
	}
	fmt.Println("转换后：", map3)

}
