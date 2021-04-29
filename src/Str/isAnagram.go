package main

import "fmt"

//判断字符串s是不是字符串t的字母异位词
func main()  {

	s := "ac"
	t := "bb"

	b := isAnagram(s, t)
	fmt.Println(b)
}
func isAnagram(s,t string) bool{

	if len(s)!= len(t) {
		return false
	}
	var c1,c2 [26]int
	for _, ch := range s {
		c1[ch-'a']++
	}
	for _, ch := range t {
		c2[ch-'a']++
	}
	//判断数组是否相等
	return c1 == c2
}
func isAnagram2(s,t string) bool  {
	if len(s) != len(t) {
		return false
	}
	cnt := map[rune]int{}
	for _, ch := range s {
		cnt[ch]++
	}
	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}
	return true
}