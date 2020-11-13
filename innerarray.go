package main

import "fmt"

func main(){
	array3 := [5]string{"a","b","c","d","e"}
	slice3 := array3[:3]
	slice4 := array3[2:5]
	
	fmt.Println(slice3, slice4)
	//array3[2] = "x"

	slice3[2] = "Queen"
	fmt.Println(array3)
	fmt.Println(slice3, slice4)
}
