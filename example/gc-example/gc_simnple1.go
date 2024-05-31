package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func longChain(depth int) {
	if depth > 0 {
		longChain(depth - 1)
	} else {
		createBigObjects()
	}
}

func createBigObjects() {
	for i := 0; i < 10000; i++ {
		bigObj := make([]byte, 1<<20) // 创建一个大对象（1MB）
		runtime.KeepAlive(bigObj)     // 防止编译器优化掉 bigObj
	}
}

func main() {
	go http.ListenAndServe(":6060", nil) // 启动 pprof HTTP 服务器

	i := 1
	for {
		fmt.Printf("触发第 %d 次垃圾回收...\n", i)
		longChain(10)           // 长链路调用，创建大量临时对象
		time.Sleep(time.Second) // 等待一段时间，让 GC 有机会运行
		i++
	}
}
