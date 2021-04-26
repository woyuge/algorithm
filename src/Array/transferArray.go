package main

import "fmt"

//旋转数组

func main()  {
	a :=[]int{1,3,4,5,6,7}
	 rotate(a,3)

}
//关键在于这段代码
func rotate(nums []int, k int)  {
	//旋转位
	var length int = len(nums)

	var temp = make([]int,length)//动态数组
	for i,v:= range nums{
		temp[i] = v
	}
	for i:=0; i<length;i++{
			nums[(i+k)%length] = temp[i]
		}
	fmt.Println(nums)
}