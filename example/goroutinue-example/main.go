package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(2)
	// 创建一个协程，它将每隔一秒打印一条消息
	go func() {
		defer wg.Done()
		for {
			fmt.Println("我是正常的协程，我在工作。")
			time.Sleep(time.Second)
		}
	}()

	// 创建另一个协程，它将立即引发崩溃
	go func() {
		fmt.Println("我是要崩溃的协程。")
		panic("崩溃了！")
		wg.Done()
	}()

	// 主协程休眠一段时间，以观察其他协程的行为
	time.Sleep(5 * time.Second)
}
