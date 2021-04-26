package main
//最长匹配数
import (
	"fmt"
    "strings"
)

func main()  {
	strs := []string{"apple","append","apolong"}
	prefix := longestCommonPrefix(strs)
	fmt.Println(prefix)
}

//查找最长通用前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) < 1{
		return ""
	}
	prefix:=strs[0]
	for _,k :=range strs{
		for strings.Index(k,prefix) !=0 {
			if len(prefix)==0 {
				return ""
			}
			prefix = prefix[:len(prefix) - 1]
		}
	}
	return prefix
}
