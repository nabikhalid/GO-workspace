package main

import "fmt"

func main(){

	i := 7
	fmt.Println(i)
	fmt.Println(&i) // memory adress, pointer to i

	// inc(i)
	// fmt.Println(i)
	// did not work 

	inc(&i) // & sends pointer
	fmt.Println(i)

}

func inc(x *int) { // * accept pointers
	*x++
}