package main

import (
	"fmt"
	"strings"
)

//验证回文串
func main()  {

	s := "abcadfsdfa"
	b := isPalindrome(s)
	fmt.Println(b)
}
//双指针算法
func isPalindrome(s string) bool {

	var sgood string

	for i:=0;i<len(s);i++{
		if isAlnum(s[i]) {
			sgood += string(s[i])
		}
	}
	n := len(sgood)
	sgood = strings.ToLower(sgood)
	for i:=0; i< n/2;i++{
		if sgood[i] != sgood[n-1-i]{
			return false
		}
	}
	return true
}
func isAlnum(ch byte) bool  {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}