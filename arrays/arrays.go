package main
import "fmt"

func main() {

	var array [5]int
	array[2] = 7
	fmt.Println(array)

	// another way to initialize an array

	barray := [5]int{1, 1, 1, 1, 7}
	fmt.Println(barray)

	// slices

	slice := []int{7, 7}
	fmt.Println(slice)
	
	slice = append(slice, 13) // returns new slice with second input added
	fmt.Println(slice)
	fmt.Println(append(slice, 13))

	

}