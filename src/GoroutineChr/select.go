package main

import (
	"fmt"
	"math/rand"
)

func main()  {
	chs :=[3] chan int{
		make(chan int,3),
		make(chan int,3),
		make(chan int,3),
	}

	index := rand.Intn(3)//随机生成0-2之间的数字
	fmt.Printf("随机索引/数值 ： %d\n", index)
	chs[index] <- index //向通道发送随机数字

	index1 := rand.Intn(3)
	fmt.Printf("随机索引/数值： %d\n", index1)
	chs[index1] <-index1

	index2 := rand.Intn(3)
	fmt.Printf("随机索引/数值： %d\n", index2)
	chs[index2] <-index2

	index3 := rand.Intn(3)
	fmt.Printf("随机索引/数值： %d\n", index3)
	chs[index3] <-index3

	for i:=0 ;i<3 ;i++ {
		select {
			case num,ok:= <- chs[0]:
				if!ok{
					break
				}
				fmt.Println("第一个条件分支被选中：chs[0]=>",num)
		case num, ok := <- chs[1]:
			if !ok {
				break
			}
			fmt.Println("第二个条件分支被选中: chs[1]=>", num)
		case num, ok := <- chs[2]:
			if !ok {
				break
			}
			fmt.Println("第三个条件分支被选中: chs[2]=>", num)
		default:
			fmt.Println("没有分支被选中")
		}
	}
}
