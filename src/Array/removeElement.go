package main
//原地删除出现在数组的元素
import "fmt"

func main()  {

	nums:=[]int{3,2,2,3,3,5}
	val :=3

	//a := []int{0, 1, 2, 3, 4}
	////删除第i个元素
	//i := 2
	//a = append(a[:i], a[i+1:]...)
	//fmt.Println(a)


	removeItem(nums,val)
}

func removeItem(nums []int, val int) []int {

	for i := 0; i < len(nums);{
		fmt.Println("=============================")
		fmt.Printf("i初始%d\n",i)
		fmt.Println(nums)
		if nums[i] == val {
			fmt.Println("命中")
			nums = append(nums[:i],nums[i+1:]...)//生成数组的切片
		}else{
			i++
		}
		fmt.Printf("i变化后%d\n",i)
	}
	fmt.Println(nums)
	return nums
}
