package main

import "fmt"

func main() {
	/*1、回顾数组的声明以及初始化
	  数组的长度是类型的一部分，也就是说 [5]int 和 [10]int 是两个不同的类型
	*/
	var arr1 [3]int
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30
	fmt.Println(arr1)

	var arr2 = [3]string{"php", "java", "golang"}
	fmt.Println(arr2)

	var arr3 = [...]string{"php", "java", "golang"}
	fmt.Println(arr3)

	var arr4 = [...]int{1: 2, 2: 3}
	fmt.Println(arr4)

	/*2、切片的声明 初始化*/

	var arr5 []int
	fmt.Printf("%v - %T - 长度:%v\n", arr5, arr5, len(arr5))

	var arr6 = []int{1, 2, 34, 45}
	fmt.Printf("%v - %T - 长度:%v\n", arr6, arr6, len(arr6))

	var arr7 = []int{1: 2, 2: 4, 5: 6}
	fmt.Printf("%v - %T - 长度:%v", arr7, arr7, len(arr7))

	var arr8 []int
	var arr9 = []int{1, 2, 34, 45}
	fmt.Println(arr8)        //[]
	fmt.Println(arr8 == nil) //true golang中申明切片以后 切片的默认值就是nil
	fmt.Println(arr9 == nil) //false

	//切片的循环遍历
	var strSlice = []string{"php", "java", "nodejs", "golang"}
	for i := 0; i < len(strSlice); i++ {
		fmt.Println(strSlice[i])
	}

	for k, v := range strSlice {
		fmt.Println(k, v)
	}

	//1、基于数组定义切片
	// a := [5]int{55, 56, 57, 58, 59}

	// b := a[:]                   //获取数组里面的所有值
	// fmt.Printf("%v-%T\n", b, b) //[55 56 57 58 59]

	// c := a[1:4]
	// fmt.Printf("%v-%T\n", c, c) //[56 57 58]

	// d := a[2:]
	// fmt.Printf("%v-%T\n", d, d) //[57, 58, 59]

	// e := a[:3]                  //表示获取第三个下标前面的数据
	// fmt.Printf("%v-%T\n", e, e) //[55, 56, 57]

	//2、基于切片定义切片

	// a := []string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
	// b := a[1:]
	// fmt.Printf("%v-%T\n", b, b) //[上海 广州 深圳 成都 重庆]

	//3、关于切片的长度和容量
	// 长度：切片的长度就是它所包含的元素个数
	// 容量：切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。

	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("长度%d 容量%d\n", len(s), cap(s)) //长度6 容量6

	a := s[2:]                                //5, 7, 11, 13
	fmt.Printf("长度%d 容量%d\n", len(a), cap(a)) //长度4 容量4

	b := s[1:3]                               //3, 5
	fmt.Printf("长度%d 容量%d\n", len(b), cap(b)) //长度2 容量5

	c := s[:3]                                //2, 3, 5
	fmt.Printf("长度%d 容量%d\n", len(c), cap(c)) //长度3 容量6

	//make()函数创建一个切片  make([]T, size, cap)
	// var sliceA = make([]int, 4, 8)
	// fmt.Println(sliceA) //[0 0 0 0]
	// fmt.Printf("%d--%d", len(sliceA), cap(sliceA))

	//修改切片内的数值
	// sliceB := []string{"php", "java", "go"}
	// sliceB[2] = "golang"
	// fmt.Println(sliceB)

	/*
		Go 语言的内建函数append()可以为切片动态添加元素，每个切片会指向一个底层数组，这个数组的容量够用就添加新增元素。
		当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行“扩容”，此时该切片指向的底层数组就会更换。
		“扩容”操作往往发生在append()函数调用时，所以我们通常都需要用原变量接收append 函数的返回值
	*/
	// var sliceA []int
	// sliceA = append(sliceA, 12)
	// sliceA = append(sliceA, 24)
	// fmt.Printf("%v - %v--%v\n", sliceA, len(sliceA), cap(sliceA)) //[12 24] - 2--2

	// sliceA = append(sliceA, 12, 23, 35, 465)
	// fmt.Printf("%v - %v--%v\n", sliceA, len(sliceA), cap(sliceA)) //[12 23 35 465] - 4--4

	//2、append方法还可以合并切片
	// sliceA := []string{"php", "java"}
	// sliceB := []string{"nodejs", "go"}
	// sliceA = append(sliceA, sliceB...)
	// fmt.Println(sliceA) //[php java nodejs go]

	//3、切片的扩容策略(了解)
	/*
		1、首先判断，如果新申请容量（cap）大于 2 倍的旧容量（old.cap），最终容量（newcap） 就是新申请的容量（cap）。
		2、否则判断，如果旧切片的长度小于 1024，则最终容量(newcap)就是旧容量(old.cap)的两 倍，即（newcap=doublecap），
		3、否则判断，如果旧切片长度大于等于 1024，则最终容量（newcap）从旧容量（old.cap） 开始循环增加原来的 1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量 （newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
		4、如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。

		需要注意的是，切片扩容还会根据切片中元素的类型不同而做不同的处理，比如 int 和 string 类型的处理方式就不一样。
	*/
	var sliceA []int
	for i := 1; i <= 10; i++ {
		sliceA = append(sliceA, i)
		fmt.Printf("%v 长度:%d 容量:%d\n", sliceA, len(sliceA), cap(sliceA))
	}
}
