package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//syncTest()
	chanTest()
}

//Go 语言提供了 sync 和 channel 两种方式支持协程(goroutine)的并发。

var wg sync.WaitGroup

func download(url string) {
	fmt.Println("start download url is", url)
	time.Sleep(time.Second)
	wg.Done() //为wg减去一个计数
}

func syncTest() {
	for i := 0; i < 3; i++ {
		wg.Add(1)                            //为wg添加一个计数
		go download("a.com" + string(i+'0')) //启动新的协成 并发执行download
	}
	wg.Wait() //等待所有的协成执行完毕
	fmt.Println("Done")
}

var ch = make(chan string, 10) // 创建大小为 10 的缓冲信道

func chanDownload(url string) {
	fmt.Println("start download url is", url)
	time.Sleep(time.Second)
	ch <- url //将URL信息发送给 chan
}

//使用 channel 信道，可以在协程之间传递消息。阻塞等待并发协程返回消息。
func chanTest() {
	for i := 0; i < 3; i++ {
		go chanDownload("a.com" + string(i+'0')) //启动新的协成 并发执行download
	}
	for i := 0; i < 3; i++ {
		msg := <-ch
		fmt.Println("finish :", msg)
	}
	fmt.Println("Done")
}
