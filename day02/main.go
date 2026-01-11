package main

import "fmt"

func main() {
	/* 声明变量
		一、变量命名规则：
	 	1、变量名称的命名：Go 语言变量由字幕、数字、下划线组成，其中首字母不能是数字
		2、Go语言中关键字和保留字不能做变量名称
	*/

	// 二、变量的定义及初始化

	//1、第一种初始化变量方式、第二种初始化以及第三种
	var username string
	username = "张三"
	fmt.Println(username)

	var sex string = "男"
	fmt.Println(sex)

	var age = "20"
	fmt.Println(age)

	// 注意：Go 语言中的变量需要声明后才能使用，同一作用域内不支持重复声明
	// var username = "李四"

	/*
		2、一次定义多个变量

				var 变量名称, 变量名称 类型

				var (
					变量名称 类型
					变量名称 类型
				)
	*/

	/*
		3、短变量声明法 在函数内部，可以使用更简略的 := 方式声明并初始化变量。

		注意：短变量只能用于声明局部变量，不能用于全局变量的声明
	*/

	//var username1 = "赵五"
	username1 := "赵6"
	fmt.Println(username1)

	// 使用短变量一次声明多个变量，并初始化变量
	a, b, c := 12, 13, 20
	fmt.Println(a, b, c)

	/*
		4、匿名变量 在使用多重赋值时，如果想要忽略某个值，可以使用匿名变量（anonymous variable）。

		匿名变量用一个下划线_表示

	*/
	var name, _ = getUserinfo()
	fmt.Println(name)
	//fmt.Println(name, _) 报错：cannot use _ as value or type
	//匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明
	var _, age1 = getUserinfo()
	fmt.Println(age1)
}

func getUserinfo() (string, int) {
	return "zhangsan", 10
}
