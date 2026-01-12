package main

import "fmt"

//接口是一个规范
type Usber interface {
	start()
	stop()
}

//如果接口里面有方法的话，必要要通过结构体或者通过自定义类型实现这个接口

type Phone struct {
	Name string
}

//手机要实现usb接口的话必须得实现usb接口中的所有方法
func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}

type Camera struct {
}

func (p Camera) start() {
	fmt.Println("相机启动")
}

func (p Camera) stop() {
	fmt.Println("相机关机")
}

func (p Camera) run() {
	fmt.Println("run")
}

//电脑
type Computer struct {
}

// var usb Usber =camera

//思想理解：电脑需要USB这个接口作为参数，才能与其他设备一起运行
func (c Computer) work(usb Usber) {
	usb.start()
	usb.stop()
}

func main() {
	p := Phone{
		Name: "华为手机",
	}
	// p.start1()

	var p1 Usber //golang中接口就是一个数据类型
	p1 = p       //表示手机实现Usb接口
	p1.stop()

	c := Camera{}
	var c1 Usber = c //表示相机实现了Usb接口
	c1.start()
	// c1.run()  //错误：type Usber has no field or method run)
	//接口里边没有定义run方法，因此你用相机实现的接口去调用run是不行的，run属于相机结构体内的方法
	c.run() //run

	var computer = Computer{}
	var phone = Phone{
		Name: "小米",
	}
	var camera = Camera{}

	computer.work(phone)

	computer.work(camera)
}
