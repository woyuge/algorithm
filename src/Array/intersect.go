package main
//取数组的交集
import (
	"fmt"
	"sort"
)
//
func main()  {
	num1 := []int {1,2,3,4,4,13}
	num2 := []int {1,2,3,9,10}

	num3 := intersect(num1, num2)
	num4 := intersect2(num1, num2)
	fmt.Println(num3)
	fmt.Println(num4)
}
/**
 */
//取数组的交集
func intersect(num1 []int, num2 []int) []int {
	m0 := map[int]int{}
	for _,v := range num1{
		m0[v] +=1
	}
	k:=0
	for _,v := range num2{
		if m0[v] > 0{
			m0[v] -= 1//跟去重有关
			num2[k] = v
			k++
		}
	}
	return num2[0:k]
}
/*
 * 双指针写法
 */
func intersect2(num1[]int, num2[]int) []int{

	sort.Ints(num1)
	sort.Ints(num2)
	length1,length2 := len(num1),len(num2)
	index1,index2 := 0,0
	var intersection []int

	for index1 < length1 && index2 < length2 {
		if num1[index1] < num2[index2]{
			index1 ++
		}else if num1[index1] > num2[index2]{
			index2 ++
		}else{
			intersection = append(intersection, num1[index1])
			index1 ++
			index2 ++
		}
	}
	return intersection
}
