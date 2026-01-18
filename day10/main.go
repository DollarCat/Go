package main

import (
	"fmt"
	"time"
)

// panic
// 函数
func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 50)
		fmt.Println("hello,world")
	}
}

// 函数
func test() {
	//这里我们可以使用defer + recover
	defer func() {
		//捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()
	//定义了一个map
	var myMap map[int]string
	myMap[0] = "golang" //error
}

func main() {
	//单向管道
	// 1、在默认情况下下，管道是双向
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 12
	m1 := <-ch1
	m2 := <-ch1
	fmt.Println(m1, m2) //10 12

	// 2、管道声明为只写
	ch2 := make(chan<- int, 2)
	ch2 <- 10
	ch2 <- 12
	// <-ch2   //receive from send-only type chan<- int

	// 3、管道声明为只读

	//ch3 := make(<-chan int, 2)
	//ch3 <- 23 //invalid operation: cannot send to receive-only channel <-chan int ch3 (variable of type <-chan int)

	// 在某些场景下我们需要同时从多个通道接收数据,这个时候就可以用到golang中给我们提供的select多路复用

	//1.定义一个管道 10个数据int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	//2.定义一个管道 5个数据string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}
	//使用select来获取channel里面的数据的时候不需要关闭channel
	for {
		select {
		case v := <-intChan:
			fmt.Printf("从 intChan 读取的数据%d\n", v)
			time.Sleep(time.Millisecond * 50)
		case v := <-stringChan:
			fmt.Printf("从 stringChan 读取的数据%v\n", v)
			time.Sleep(time.Millisecond * 50)
		default:
			fmt.Printf("数据获取完毕\n")
			return //注意退出...
		}
	}

	//panic
	go sayHello()
	go test()

	//防止主进程退出这里使用time.Sleep演示，搭建也可以用sync.WaitGroup
	time.Sleep(time.Second)
}
