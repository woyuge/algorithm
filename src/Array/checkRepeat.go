package main

//存在重复元素,用求字典的方式去判断
func main()  {

	a :=[]int{1,2,3,4,5,6,6,6,6,7,7}
	//b:= containsDuplicate(a)

	c:=containsDuplicateHaset(a)
	println(c)
}

//哈希表算法
func containsDuplicateHaset(nums []int) bool {
	m0 := map[int]int{}
	for _,v := range nums{
		if _,ok:= m0[v];ok {
			return true
		}
		m0[v]++
	}
	return false
}

//手写map，算法不算太好
func containsDuplicate(nums []int) bool {

	m0 := map[int]int{}
	for _,v := range nums{
		m0[v]+=1
	}

	for _,v := range m0{
		if v>1 {
			return true
		}
	}
	return false
}