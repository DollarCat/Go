package main

import "fmt"

func main() {
	// 1、make创建map类型的数据

	var userinfo = make(map[string]string)
	userinfo["username"] = "张三"
	userinfo["age"] = "20"
	userinfo["sex"] = "男"
	fmt.Println(userinfo)
	fmt.Println(userinfo["username"])

	//2、map 也支持在声明的时候填充元素
	var userinfo1 = map[string]string{
		"username": "张三",
		"age":      "20",
		"sex":      "男",
	}
	fmt.Println(userinfo1)

	// 3、第三种创建map类型数据的方法

	userinfo2 := map[string]string{
		"username": "张三",
		"age":      "20",
		"sex":      "男",
	}
	fmt.Println(userinfo2)

	//for range循环遍历map类型的数据
	for k, v := range userinfo {
		fmt.Printf("key:%v value:%v\n", k, v)
	}

	//map类型数据的curd
	//1、创建 修改map类型的数据
	// var userinfo = make(map[string]string)
	// userinfo["username"] = "张三"
	// userinfo["age"] = "20"
	// fmt.Println(userinfo)

	//2、创建 修改map类型的数据

	// var userinfo = map[string]string{
	// 	"username": "张三",
	// 	"age":      "20",
	// }

	// userinfo["username"] = "李四"
	// fmt.Println(userinfo)

	//3、获取 查找map类型的数据

	var userinfo3 = map[string]string{
		"username": "张三",
		"age":      "20",
	}
	fmt.Println(userinfo3["username"]) //获取

	// v, ok := userinfo3["age"]
	// fmt.Println(v, ok) //20 true

	v, ok := userinfo3["xxxxx"]
	fmt.Println(v, ok) // 空 和 false

	//4、删除map数据里面的key以及对于的值

	var userinfo4 = map[string]string{
		"username": "张三",
		"age":      "20",
		"sex":      "男",
		"height":   "180cm",
	}
	fmt.Println(userinfo4)

	delete(userinfo4, "sex")
	fmt.Println(userinfo4)

	/*
		值类型 ：改变变量副本值的时候，不会改变变量本身的值  (基本数据类型、数组)
		引用类型：改变变量副本值的时候，会改变变量本身的值 （切片、map）
	*/

	//map类型也是引用数据类型
	var userinfo5 = make(map[string]string)
	userinfo5["username"] = "张三"
	userinfo5["age"] = "20"
	userinfo6 := userinfo5

	userinfo6["username"] = "李四"
	fmt.Println(userinfo5)
	fmt.Println(userinfo6)

	//如果我们想在map对象中存放一系列的属性的时候，我们就可以把map类型的值定义成切片
	// var userinfo = make(map[string]string)
	// userinfo["username"] = "张三"
	// userinfo["hobby"] = "睡觉"

	var userinfo7 = make(map[string][]string)

	userinfo7["hobby"] = []string{
		"吃饭",
		"睡觉",
		"敲代码",
	}

	userinfo7["work"] = []string{
		"php",
		"golang",
		"前端",
	}

	// fmt.Println(userinfo)

	/*
		_ 空白标识符是 Go 语言中一个非常实用的特性，它允许你：
		- 忽略不需要的返回值
		- 避免未使用变量的编译错误
		- 使代码意图更清晰
	*/
	for _, v := range userinfo7 {
		// fmt.Println(k, v)
		for _, value := range v {
			fmt.Println(value)
		}
	}
}
