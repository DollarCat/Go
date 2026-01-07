package main

import "fmt"

/*
	函数定义：
	func 函数名(参数)(返回值){
		函数体
	}
*/

//函数参数的简写
func subFn(x, y int) int {
	sub := x - y
	return sub
}

//函数的可变参数，可变参数是指函数的参数数量不固定。Go 语言中的可变参数通过在参数名后加...来标识
func sumFn1(x ...int) int {

	// fmt.Printf("%v--%T", x, x) //[12 34 45 46]--[]int

	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

//函数返回值：
//return 关键词一次可以返回多个值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

//返回值命名: 函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过 return 关键字返回。
func calc1(x, y int) (sum int, sub int) {
	fmt.Println(sum, sub)
	sum = x + y
	sub = x - y
	fmt.Println(sum, sub)
	return
}

func main() {
	sum := sumFn1(100, 1, 2, 3, 4)
	fmt.Println(sum)

	a, b := calc(10, 2)
	fmt.Println(a, b)
}
