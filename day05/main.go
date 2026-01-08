package main

import "fmt"

func fn1(x int) {
	x = 10
}
func fn2(x *int) {
	*x = 40
}

func main() {
	//每一个变量都有自己的内存地址
	var a = 10
	var p = &a //p指针变量   p的类型 *int（指针类型）
	fmt.Printf("a的值：%v a的类型%T a的地址%p\n", a, a, &a)
	fmt.Printf("p的值：%v p的类型%T p的地址%p", p, p, &p)

	c := 10
	d := &c //d指针变量  类型 *int

	// *d ：表示取出p这个变量对应的内存地址的值
	fmt.Println(d)  //c的地址  0xc0000100a8
	fmt.Println(*d) //*d 表示打印a对应的值   10

	*d = 30 //改变d这个变量对应的内存地址的值

	fmt.Println(c) //30

	//指针传值
	var e = 5
	fn1(e)
	fmt.Println(e) //5
	fn2(&e)
	fmt.Println(e) //40

	// 实际想开发中 new 函数不太常用，使用 new 函数得到的是一个指类型针，并且该指针对应的值为该类型的零值
	var f = new(int) //a是一个指针变量 类型是 *int的指针类型 指针变量对应的值是0

	fmt.Printf("值：%v 类型:%T 指针变量对应的值：%v", f, f, *f) //值：0xc0000a0090 类型:*int 指针变量对应的值：0

	/*
		错误的写法：没有分配内存空间
		var a *int //指针也是引用数据类型
		*a = 100
		fmt.Println(*a)
	*/

	//正确的写法：new方法给指针变量分配存储空间
	var g *int
	g = new(int)
	*g = 100
	fmt.Println(*g)

	//make函数
	/*
		错误写法：没有分配内存空间
		var userinfo map[string]string
		userinfo["username"] = "张三"
		fmt.Println(userinfo)
	*/

	var userinfo = make(map[string]string)
	userinfo["username"] = "张三"
	fmt.Println(userinfo)

	var h = make([]int, 4, 4)
	h[0] = 1
	fmt.Println(h)
}
