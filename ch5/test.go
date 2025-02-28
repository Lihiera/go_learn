package main

import (
	"fmt"
	"time"
)

// 模拟一个耗时的异步操作，假设它在 3 秒钟后完成
func downloadFile(url string, callback func(result string)) {
	fmt.Println("开始下载文件:", url)

	// 模拟下载操作，异步执行
	go func() {
		// 假设下载过程需要 3 秒
		time.Sleep(3 * time.Second)

		// 下载完成后，调用回调函数
		callback("下载完成: " + url)
	}()
}

func main() {
	// 使用回调函数处理异步操作完成后的逻辑
	downloadFile("https://example.com/file", func(result string) {
		fmt.Println(result) // 输出回调信息
	})

	// 主程序继续做其他事情，不会被下载操作阻塞
	fmt.Println("继续执行其他任务...")

	// 防止主程序提前结束，给异步操作时间完成
	time.Sleep(4 * time.Second)
}
