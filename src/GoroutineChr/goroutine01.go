package main

import (
	"fmt"
)

func add(a,b int) {
	var c= a+b
	fmt.Printf("%d + %d= %d\n",a, b, c)
}
func main()  {

	for i := 0; i < 10; i++ {
		go add(1, i)
	}
	fmt.Println("a")
	//time.Sleep(1e9)
}
