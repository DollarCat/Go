package main

import "fmt"

//自定义类型
type myInt int

// type myFn func(int, int) int

//类型别名
type myFloat = float64

//结构体定义
//注意：结构体首字母可以大写也可以小写，大写表示这个结构体是公有的，在其他的包里面 可以使用。小写表示这个结构体是私有的，只有这个包里面才能使用
type Person struct {
	name string
	age  int
	sex  string
}

func main() {
	var a myInt = 10

	fmt.Printf("%v %T\n", a, a) //10 main.myInt

	var b myFloat = 12.3
	fmt.Printf("%v %T\n", b, b) //12.3 float64

	var p1 Person //实例化Person结构体
	p1.name = "张三"
	p1.sex = "男"
	p1.age = 20
	fmt.Printf("值:%v 类型:%T\n", p1, p1) //值:{张三 20 男} 类型:main.Person
	//%#v会把类型也打印出来，比较常用
	fmt.Printf("值:%#v 类型:%T\n", p1, p1) //值:main.Person{name:"张三", age:20, sex:"男"} 类型:main.Person

	//注意：在 Golang 中支持对结构体指针直接使用.来访问结构体的成员。p2.name = "张三" 其实在底层是(*p2).name = "张三"
	var p2 = new(Person)
	p2.name = "李四"
	p2.age = 20
	p2.sex = "男"
	(*p2).name = "王五"
	fmt.Printf("值:%#v 类型:%T\n", p2, p2) //值:&main.Person{name:"王五", age:20, sex:"男"} 类型:*main.Person

	//结构体有5、6种初始化方式，就不多写了。。。

	/*
		值类型 ： 改变变量副本值的时候，不会改变变量本身的值 (数组、基本数据类型、结构体)
		引用类型：改变变量副本值的时候，会改变变量本身的值  （切片、map）
	*/

	var p3 = Person{
		name: "哈哈",
		age:  20,
		sex:  "男",
	}

	p4 := p3
	p4.name = "李四"
	fmt.Printf("%#v\n", p3) //main.Person{Name:"哈哈", Age:20, Sex:"男"}
	fmt.Printf("%#v\n", p4) //main.Person{Name:"李四", Age:20, Sex:"男"}
}
