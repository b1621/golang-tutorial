package main 

import "fmt"

func concat(s1 string, s2 string) string{
	return s1 + s2
}
//  if both arguments are  same types only need to declare after the last one
func add(x, y int) int{
	return x + y
}
func main(){

	fmt.Println(concat("lame,", "happy birthday"))
}