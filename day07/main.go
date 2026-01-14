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

/*
1、手机要实现usb接口的话必须得实现usb接口中的所有方法

2、结构体值接收者实现接口：
值接收者： 如果结构体中的方法是值接收者，那么实例化后的结构体值类型和结构体指针类型都可以赋值给接口变量

*/
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

//golang中空接口也可以直接当做类型来使用，可以表示任意类型

//空接口作为函数的参数
func show(a interface{}) {
	fmt.Printf("值:%v 类型:%T\n", a, a)
}

/*类型断言：用于判断空接口中值的类型
定义一个方法，可以传入任意数据类型，然后根据不同的类型实现不同的功能
原型：x.(T)
x : 表示类型为 interface{}的变量
• T : 表示断言 x 可能是的类型
该语法返回两个参数，第一个参数是 x 转化为 T 类型后的变量，第二个值是一个布尔值，若
为 true 则表示断言成功，为 false 则表示断言失败
*/
func MyPrint1(x interface{}) {
	if _, ok := x.(string); ok {
		fmt.Println("string类型")
	} else if _, ok := x.(int); ok {
		fmt.Println("int类型")
	} else if _, ok := x.(bool); ok {
		fmt.Println("bool类型")
	}
}

//x.(type) 判断一个变量的类型  这个语句只能用在switch语句里面
func MyPrint2(x interface{}) {

	switch x.(type) {
	case int:
		fmt.Println("int类型")
	case string:
		fmt.Println("string类型")
	case bool:
		fmt.Println("bool类型")
	default:
		fmt.Println("传入错误...")
	}
}

/*
结构体指针接收者实现接口：

指针接收者： 如果结构体中的方法是指针接收者，那么实例化后结构体指针类型都可以赋值给接口变量， 结构体值类型没法赋值给接口变量。

*/
type Usber1 interface {
	start1()
	stop1()
}

func (p *Phone) start1() { //指针接收者
	fmt.Println(p.Name, "启动")
}

func (p *Phone) stop1() {
	fmt.Println(p.Name, "关机")
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

	type A interface{} //空接口  表示没有任何约束  任意的类型都可以实现空接口

	var a A
	var str = "你好golang"
	a = str                          //让字符串实现A这个接口
	fmt.Printf("值:%v 类型:%T\n", a, a) //值:你好golang 类型:string

	var num = 20
	a = num //表示让int类型实现A这个接口
	fmt.Printf("值:%v 类型:%T\n", a, a)

	var flag = true
	a = flag //表示让bool类型实现A这个接口
	fmt.Printf("值:%v 类型:%T\n", a, a)

	show(20)
	show("你好golang\n")
	slice := []int{1, 2, 34, 4}
	show(slice)

	//类型断言
	var aa interface{}
	aa = "你好golang"
	v, ok := aa.(string)
	if ok {
		fmt.Println("a就是一个string类型,值是：", v)
	} else {
		fmt.Println("断言失败")
	}

	MyPrint1("你好golang")
	MyPrint1(true)

	MyPrint2("你好golang")
	MyPrint2(true)

	// 结构体值接收者例化后的结构体值类型和结构体指针类型都可以赋值给接口变量
	fmt.Println("\n结构体值接收者例化后的结构体值类型和结构体指针类型都可以赋值给接口变量\n")
	var p2 = Phone{
		Name: "小米手机",
	}
	var p3 Usber = p2 //表示让Phone实现Usb的接口
	p3.start()

	var p4 = &Phone{
		Name: "苹果手机",
	}
	var p5 Usber = p4 //表示让Phone实现Usb的接口
	p5.start()

	//指针接收者
	fmt.Println("\n指针接收者\n")
	var phone1 = &Phone{
		Name: "苹果15pro max",
	}
	var p6 Usber1 = phone1
	p6.start1()
	p6.stop1()
}
