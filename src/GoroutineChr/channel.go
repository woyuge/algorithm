package main

import (
	"fmt"
	"time"
)

func add(a, b int, ch chan int)  {
	c := a + b
	fmt.Printf("%d + %d = %d\n",a, b, c)
	ch <- 1
}
func test(ch chan int)  {
	for i := 0; i< 10;i++{
		ch <- i
	}
	close(ch)
}

func main()  {
	start:= time.Now()
	ch := make(chan int)
	go test(ch)
	for i := range ch {
		fmt.Println("接收到的数据:", i)
	}

	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
