package main

import "fmt"

func main() {
	// if语句写法
	flag := true
	if flag {
		fmt.Println("flag=true")
	}

	// age := 30
	// if age > 20 {
	// 	fmt.Println("成年人")
	// }

	// 1、这样写的age属于局部变量 因此可以重复定义age这个变量
	// 2、{必须挨着if和else
	if age := 35; age > 30 {
		fmt.Println(age)
	}

	// for 语句
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// for 循环的初始语句和结束语句都可以省略，例如
	i := 1
	for i <= 5 { //注意：Go 语言中是没有 while 语句的，我们可以通过 for 代替
		fmt.Println(i)
		i++
	}

	/*
		for 无限循环
		for {
			循环体语句
		}
	*/

	// switch语句
	// Go 语言规定每个 switch 只能有一个 default 分支
	extname := ".a"
	switch extname {
	case ".html":
		fmt.Println("text/html")
		break
	case ".css":
		fmt.Println("text/css")
		break
	case ".js":
		fmt.Println("text/javascript")
		break
	default:
		fmt.Println("格式错误")
		break
	}

	// 2、switch case的另一种写法

	switch extname := ".html"; extname { //同样是局部变量
	case ".html":
		fmt.Println("text/html")
		break
	case ".css":
		fmt.Println("text/css")
		break
	case ".js":
		fmt.Println("text/javascript")
		break
	default:
		fmt.Println("找不到此后缀")
	}
	//fmt.Println(extname) //undefined: extname

	//3、一个分支可以有多个值，多个 case 值中间使用英文逗号分隔
	// 判断一个数是不是偶数

	var n = 8
	switch n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
		break //golang中break可以写也可以不写
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
		break
	}

	//5、 switch 的穿透 fallthrought

	//fallthrough`语法可以执行满足条件的 case 的下一个 case，是为了兼容 C 语言中的 case 设计 的

	var age = 30
	switch {
	case age < 24:
		fmt.Println("好好学习")
	case age >= 24 && age <= 60:
		fmt.Println("好好赚钱")
		fallthrough
	case age > 60:
		fmt.Println("注意身体")
	default:
		fmt.Println("输入错误")
	}
}
