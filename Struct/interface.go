package main

import (
	"fmt"
)

//USB : usb接口
type USB interface {
	Name() string
	Connect //组合接口，USB接口可以转换为Connect接口
}

//Connect : connect接口
type Connect interface {
	Connect()
}

//PhoneDevice : 手机类型
type PhoneDevice struct {
	name string
}

//Name : 实现USB接口的方法
func (p PhoneDevice) Name() string {
	return p.name
}

//Connect : 实现Connect接口方法
func (p PhoneDevice) Connect() {
	fmt.Println("Connect：", p.name)
}

func main() {
	var p USB                 // 使用的是USB接口
	p = PhoneDevice{"XiaoMI"} //因为PhoneDevice 实现了USB接口，所以可以将它赋值给USB

	a := PhoneDevice{"IPhone"} //直接赋值
	fmt.Println(p.Name())
	p.Connect()
	DisConnect(a) //调用方法，传入实现USB接口的类型

	//将USB接口转换为Connect接口
	u := PhoneDevice{"Huawei"}

	//var u USB
	//var c Connect
	//u = USB(c) //connect接口无法转换为USB接口 因为connect接口没有Name()方法，编译时就会报错
	var c Connect
	c = Connect(u) //USB接口转换为connect接口
	c.Connect()

}

//DisConnect : 必须传入USB
func DisConnect(usb interface{}) {
	//使用类型判断
	/*if p, ok := usb.(PhoneDevice); ok {
		fmt.Println("Disconnect：", p.name)
	}*/
	//类型断言 type swich
	switch v := usb.(type) {
	case PhoneDevice:
		fmt.Println("Disconnect：", v.name)
	default:
		fmt.Println("Unknow Device")
	}
}
