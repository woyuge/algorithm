package main

import (
	"fmt"
	"math"
)

func main()  {
	var a int =123
	i:=reverse(a)

	fmt.Println(i)
}
//123

/***
	pop =3, x=12 ret=3
	pop =2, x=1 ret=32
	pop =1, x=0 ret=321
 */

func reverse(x int) int {

	ret := 0
	for x!=0 {
		pop := x % 10 //每次取末尾数字
		x = x/10
		ret = ret*10 + pop
		if ret < math.MinInt32 || ret > math.MaxInt32 {
			return 0
		}
	}

	return ret
}
