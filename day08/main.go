package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 在主线程(可以理解成进程)中，开启一个goroutine, 该协程每隔50毫秒输出 "你好golang"
// 在主线程中也每隔50毫输出"你好golang", 输出10次后，退出程序
// 要求主线程和goroutine同时执行
func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() 你好golang")
		time.Sleep(time.Millisecond * 100)
	}
}

// 主线程退出后所有的协程无论有没有执行完毕都会退出，所以我们在主进程中可以通过WaitGroup等待协程执行完毕
var wg sync.WaitGroup

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("test1() 你好golang-", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() //协程计数器-1
}

func test2() {
	for i := 0; i < 10; i++ {
		fmt.Println("test2() 你好golang-", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() //协程计数器-1
}

// 多协程案例
func test3(num int) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("协程(%v)打印的第%v条数据\n", num, i)
		time.Sleep(time.Microsecond * 100)
	}
}

// 统计1-120000的数字中那些是素数？goroutine  for循环实现
func test4(n int) {
	for num := (n-1)*30000 + 1; num < n*30000; num++ {
		if num > 1 {
			var flag = true
			for i := 2; i < num; i++ {
				if num%i == 0 {
					flag = false
					break
				}
			}
			if flag {
				// fmt.Println(num, "是素数")
			}
		}
	}
	wg.Done()
}

func main() {
	// go test() //表示开启一个协程
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("main() 你好golang")
	// 	time.Sleep(time.Millisecond * 100)
	// }

	// wg.Add(1)  //协程计数器+1
	// go test1() //表示开启一个协程
	// wg.Add(1)  //协程计数器+1
	// go test2() //表示开启一个协程

	// wg.Wait() //等待协程执行完毕...
	// fmt.Println("主线程退出...")

	//获取当前计算机上面的Cup个数
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)

	//可以自己设置使用多个cpu
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")

	// 多协程案例
	// 10个goroutine是并发执行的
	// for i := 1; i <= 6; i++ {
	// 	wg.Add(1)
	// 	go test3(i)
	// }
	// wg.Wait()
	// fmt.Println("关闭主线程...")

	//统计1-120000的数字中那些是素数？for循环实现
	// start := time.Now().Unix()
	// for num := 2; num < 120000; num++ {
	// 	var flag = true
	// 	for i := 2; i < num; i++ {
	// 		if num%i == 0 {
	// 			flag = false
	// 			break
	// 		}
	// 	}
	// 	if flag {
	// 		fmt.Println(num, "是素数")
	// 	}
	// }
	// end := time.Now().Unix()

	// fmt.Println(end - start) //6秒

	//统计1-120000的数字中那些是素数？goroutine  for循环实现
	start := time.Now().Unix()
	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go test4(i)
	}
	wg.Wait()
	fmt.Println("执行完毕")
	end := time.Now().Unix()
	fmt.Println(end - start) //2毫秒
}
