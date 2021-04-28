package main

import (
	"fmt"
	"time"
)

func main()  {

	ch := make(chan int, 1)

	timeout := make(chan bool,1)

	go func() {
		time.Sleep(1e9)
		timeout <- true
	}()

	select {
		case <-ch:
			fmt.Println("接收到 ch 通道数据")
		case <-timeout:
			fmt.Println("超时1s，程序退出")
	}
}
