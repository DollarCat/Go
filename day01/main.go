package main

import "fmt"

func main() {
	fmt.Print("A")
	fmt.Print("B")
	fmt.Print("C")
	fmt.Print("\n")

	fmt.Println("A")
	fmt.Println("B")
	fmt.Println("C")

	fmt.Print("A", "B", "C")
	fmt.Print("\n")
	fmt.Println("A", "B", "C")

	var a int = 10 //go 中定义变量必须要使用，否则会报错
	fmt.Printf("%v, a的类型是：%T\n", a, a)

	//类型推导方式定义变量
	b := 100
	fmt.Printf("%v, b的类型是：%T", b, b)

}
