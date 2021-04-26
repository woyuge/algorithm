package main

import "fmt"

func main()  {

	s:=[]byte{'a','c','d','e','f'}
	reverseString(s)
}


func reverseString(s []byte)  {

	strLen := len(s)
	right := strLen-1

	for i := 0; i < strLen/2; i++ {
		temp := s[i]
		s[i] = s[right -i]
		s[right-i]= temp

	}
	fmt.Println(string(s))
}
