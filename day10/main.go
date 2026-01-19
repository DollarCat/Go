package main

import (
	"fmt"
	"sync"
	"time"
)

/* panic
panic 是Go语言中的一种内置机制，用于处理程序运行时的严重错误。
当程序遇到无法处理的异常情况时，会触发panic，导致程序立即终止当前函数的执行，
并开始回溯 （unwinding）调用栈，依次执行各层函数的defer语句，然后终止程序并输出错误信息
*/
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
		/*捕获test抛出的panic
		- 当程序发生panic时，recover()会捕获panic的值，并停止panic的传播（即不会继续回溯调用栈），使程序可以继续执行后续代码。
		- 返回值 ：如果捕获到panic，返回panic的值；否则返回nil。
		- 限制 ：recover()只有在defer函数中调用时才有效，在其他地方调用会返回nil
		*/
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()
	//定义了一个map
	//map变量声明后默认为 nil ，必须使用 make() 函数初始化后才能使用。当对nil map进行写操作时，Go运行时会触发panic
	var myMap map[int]string
	myMap[0] = "golang" //error
}

// lock
var count = 0
var wg sync.WaitGroup
var mutex sync.Mutex

func test_lock() {
	mutex.Lock()
	count++
	fmt.Println("the count is : ", count)
	time.Sleep(time.Microsecond)
	mutex.Unlock()
	wg.Done()
}

// mutex
var m = make(map[int]int, 0)

func test_mutex(num int) {
	mutex.Lock()
	var sum = 1
	for i := 1; i <= num; i++ {
		sum *= i
	}
	m[num] = sum
	// fmt.Println(m[num])
	fmt.Printf("key=%v value=%v\n", num, sum)
	time.Sleep(time.Millisecond)
	mutex.Unlock()
	wg.Done()
}

// 读写锁 优先级 ：写锁请求通常会优先于读锁请求
// 写的方法
var mutex_rw sync.RWMutex

func write() {
	mutex_rw.Lock()
	fmt.Println("执行写操作")
	time.Sleep(time.Second * 2)
	mutex_rw.Unlock()
	wg.Done()
}

// 读的方法
func read() {
	mutex_rw.RLock()
	fmt.Println("---执行读操作")
	time.Sleep(time.Second * 2)
	mutex_rw.RUnlock()
	wg.Done()
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
	// for {
	// 	select {
	// 	case v := <-intChan:
	// 		fmt.Printf("从 intChan 读取的数据%d\n", v)
	// 		time.Sleep(time.Millisecond * 50)
	// 	case v := <-stringChan:
	// 		fmt.Printf("从 stringChan 读取的数据%v\n", v)
	// 		time.Sleep(time.Millisecond * 50)
	// 	default:
	// 		fmt.Printf("数据获取完毕\n")
	// 		return //注意退出...
	// 	}
	// }

	//panic sayHello会正常打印出全部的hello,world
	// go sayHello()
	// go test()

	//防止主进程退出这里使用time.Sleep演示，搭建也可以用sync.WaitGroup
	//time.Sleep(time.Second)

	//lock
	// for r := 0; r < 20; r++ {
	// 	wg.Add(1)
	// 	go test_lock()
	// }
	// wg.Wait()

	//mutex
	// for r := 0; r < 40; r++ {
	// 	wg.Add(1)
	// 	go test_mutex(r)
	// }
	// wg.Wait()

	/* 开启10个协程执行写操作
	   写锁的排他性：10个写协程必须依次获得写锁，串行执行，所以每个写操作需要等待前一个写操作完成（约2秒）
	*/
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	/* 开启10个协程执行读操作
	   读锁的共享性 ：10个读协程可以同时获得读锁，并发执行，所以看起来像是同时完成
	*/
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
}
