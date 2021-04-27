package main

import (
	"fmt"
)

func main()  {

	var s string
	s = "e"
	i := firstUniqChar(s)
	fmt.Println(i)
}

/**
优势算法
 */
func firstUniqChar(s string) int{
		cnt := [26]int{}
		for _, ch := range s {
			cnt[ch-'a']++
		}
		fmt.Println(cnt)
		for i, ch := range s {
			if cnt[ch-'a'] == 1 {
			return i
		}
	}
		return -1
}

/**
劣势算法
 */
func firstUniqChar2(s string) int {
	var i int
	i= -1
	m0 := map[string]int{}
	for _,v:= range s{
		m0[string(v)] +=1
	}

re:
	for key,v:= range s{
		for ii,kk := range m0{
			if string(v) == ii && kk == 1 {
				i = key
				break re
			}
		}
	}
	return i
}