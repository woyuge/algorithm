package main
//最大收益，解题思路，最长波段
import "fmt"

func main() {
	arr := []int{1,4,9,3,5,7}
	maxNum := maxProfit2(arr)
	fmt.Printf("合计为:%d",maxNum)
}

func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][2]int, len(prices))

	dp[0][0] = 0
	dp[0][1] = -prices[0]

	fmt.Println(dp)

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0],dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][0]-prices[i],dp[i-1][1])
	}
	fmt.Println(dp)
	return dp[len(prices)-1][0]
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}



/**
	计算波段的最大收益
	解题思路，求，上升波段的累加，怎么知道两个波段的在上山阶段？？？
 */
func maxProfit(arr []int) int {
	var max int

	max = 0
	//考虑下指针？如果前面一个大于之前一个，就后移，双向后移，找到第一个买点
	index1, offerset:= 0,0
	for i:=1;i<len(arr);i++ {
		//设置边界
		if index1 <len(arr) {
			//前一个小于后一个，那么设定区间
			if arr[index1] < arr[index1+1] {
				fmt.Printf("前值为%d，后值为%d", arr[index1], arr[index1+1])
				fmt.Println()
				//指针向后移动
				index1 ++
				offerset ++ //移动了几次
			}else{
				max += arr[index1] - arr[index1 - offerset]
				fmt.Printf("索引%d的值为%d，前值为%d，并且移动了%d次",index1,arr[index1],arr[index1 - offerset], offerset)
				fmt.Println()
				offerset = 0 //位移重置为0,进行下一次
				index1 ++
			}
		}
		fmt.Println(index1, offerset)
	}
	if offerset >0 {
		max += arr[index1] - arr[index1 - offerset]
	}
	return max
}
