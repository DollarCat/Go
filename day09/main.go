package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func putNum(intChan chan int) {
	for i := 2; i < 10; i++ {
		intChan <- i
		fmt.Println(i)
	}
	close(intChan)
	wg.Done()
}

type Cat struct {
	Name string
	Age  int
}

/*
需求：使用goroutine和channel协同工作案例
1、开启一个WriteData的的协程给向管道inChan中写入100条数据
2、开启一个ReadData的协程读取inChan中写入的数据
3、注意：WriteData和ReadData同时操作一个管道
4、主线程必须等待操作完成后才可以退出

goroutine结合Channel使用的简单demo,定义两个方法，一个方法给管道里面写数据，一个给管道里面读取数据。要求同步进行。
*/

// 写数据
func fn1(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("【写入】数据%v成功\n", i)
		time.Sleep(time.Millisecond * 5000) //5秒
	}
	close(ch)
	wg.Done()
}
func fn2(ch chan int) {
	for v := range ch {
		fmt.Printf("【读取】数据%v成功\n", v)
		time.Sleep(time.Millisecond * 10)
	}
	wg.Done()
}

// 向 intChan放入 1-120000个数，具体设计参考附件goroutine&channel找素数设计.png
// 素数：就是除了 1 和它本身不能被其他数整除的数
func putNum1(intChan chan int) {
	for i := 2; i < 120000; i++ {
		intChan <- i
	}
	close(intChan)
	wg.Done()
}

// 从 intChan取出数据，并判断是否为素数,如果是，就把得到的素数放在primeChan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for num := range intChan {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num //num是素数
		}
	}
	//要关闭 primeChan
	// close(primeChan) //如果一个channel关闭了就没法给这个channel发送数据了
	//什么时候关闭primeChan?

	//给exitChan里面放入一条数据
	exitChan <- true
	wg.Done()
}

// printPrime打印素数的方法
func printPrime(primeChan chan int) {
	for v := range primeChan {
		fmt.Println(v)
	}
	wg.Done()
}

func main() {
	//1、创建channel
	ch := make(chan int, 3)

	//2、给管道里面存储数据
	ch <- 10
	ch <- 21
	ch <- 32

	//3、获取管道里面的内容
	a := <-ch
	fmt.Println(a)

	<-ch //从管道里面取值   //21
	c := <-ch
	fmt.Println(c) //32
	ch <- 56
	ch <- 66
	//4、打印管道的长度和容量
	fmt.Printf("值：%v 容量：%v 长度%v\n", ch, cap(ch), len(ch)) //值：0xc0000d0080 容量：3 长度2

	// 5、管道的类型（引用数据类型）
	ch1 := make(chan int, 4)

	ch1 <- 34
	ch1 <- 54
	ch1 <- 64

	ch2 := ch1
	ch2 <- 69
	<-ch1
	<-ch1
	<-ch1
	d := <-ch1
	fmt.Println(d) //69

	//8、管道阻塞（又叫无缓冲管道）

	// ch6 := make(chan int, 1)
	// ch6 <- 34
	// ch6 <- 64 //all goroutines are asleep - deadlock!

	// 在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock
	// ch7 := make(chan string, 2)
	// ch7 <- "数据1"
	// ch7 <- "数据2"
	// m1 := <-ch7
	// m2 := <-ch7
	// m3 := <-ch7
	// fmt.Println(m1, m2, m3) //fatal error: all goroutines are asleep - deadlock!

	//正确的写法
	ch8 := make(chan int, 1) //只要管道的容量大于零，那么该管道就是有缓冲的管道，管道的容量表示管道中能存放元素的数量
	ch8 <- 34
	<-ch8
	ch8 <- 67
	<-ch8
	ch8 <- 78
	m4 := <-ch8
	fmt.Println(m4)

	//2、使用for range遍历通道，当通道被关闭的时候就会退出for range,如果没有关闭管道就会报个错误fatal error: all goroutines are asleep - deadlock!
	var ch9 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch9 <- i
	}
	close(ch9) //关闭管道
	//close(ch9) //关闭一个已经关闭的管道会导致 panic。

	//for range循环遍历管道的值  ,注意：管道没有key
	for v := range ch9 {
		fmt.Println(v)
	}

	//2、通过for循环遍历管道的时候管道可以不关闭，不会报错
	var ch10 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch10 <- i
	}

	for j := 0; j < 10; j++ {
		fmt.Println(<-ch10)
	}

	//定义一个存放任意数据类型的管道 3个数据
	allChan := make(chan interface{}, 3)

	allChan <- 10
	allChan <- "tom jack"
	cat := Cat{"小花猫", 4} // cat 的具体类型是 Cat
	allChan <- cat       // 放入管道时，具体类型信息被"擦除"，只剩下 interface{}

	//我们希望获得到管道中的第三个元素，则先将前2个推出
	<-allChan
	<-allChan

	newCat := <-allChan //从管道中取出的Cat是什么? 取出来时，编译器只知道是 interface{}
	fmt.Printf("newCat=%T , newCat=%v\n", newCat, newCat)
	// //下面的写法是错误的!编译不通过, newCat 的类型是 interface{}，interface{} 类型没有 Name 字段
	// //fmt.Printf("newCat.Name=%v", newCat.Name)
	//使用类型断言，类型断言会知道变量实际的类型，所以得出aa为Cat类型
	// aa := newCat.(Cat)
	// fmt.Printf("aa.Name=%v", aa.Name)

	// fmt.Printf("aa=%T , aa=%v\n", aa, aa)

	// var intChan = make(chan int, 1000)
	// wg.Add(1)
	// go putNum(intChan)
	// wg.Wait()

	// fmt.Println("\n=======================需求：使用goroutine和channel协同工作案例: \n")
	// var ch11 = make(chan int, 10)

	/*
		- fn1写入第一个数据 ：写入1，然后沉睡5秒
		- fn2立即读取 ：读取到1，沉睡10毫秒
		- fn2等待 ：由于管道为空，fn2在 range 处阻塞等待
		- fn1苏醒 ：5秒后写入第二个数据2，再次沉睡5秒
		- fn2读取 ：立即读取到2，沉睡10毫秒后继续等待
		- 重复此模式 直到fn1写入完10个数据并关闭管道
	*/
	// wg.Add(1)
	// go fn1(ch11)
	// wg.Add(1)
	// go fn2(ch11)

	// wg.Wait()
	// fmt.Println("退出...")

	// 向 intChan放入 1-120000个数，具体设计参考附件goroutine&channel找素数设计.png
	start := time.Now().Unix()

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 50000)
	exitChan := make(chan bool, 16) //标识primeChan close

	//存放数字的协程
	wg.Add(1)
	go putNum1(intChan)

	//统计素数的协程
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go primeNum(intChan, primeChan, exitChan)
	}

	//打印素数的协程
	wg.Add(1)
	go printPrime(primeChan)

	//判断exitChan是否存满值
	wg.Add(1)
	go func() {
		for i := 0; i < 16; i++ {
			<-exitChan // 等待16个完成信号
		}
		// 所有素数判断协程完成后关闭
		close(primeChan)
		wg.Done()
	}()

	wg.Wait()

	end := time.Now().Unix()
	fmt.Println("执行完毕....", end-start, "毫秒")
}
