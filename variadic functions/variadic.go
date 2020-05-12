package main

import "fmt"

// take an arbitrary number of arguements as inputs, println is a variadic function

func multiply(nums ...int) int { // nums is the input
	total := 1

	for _, num := range nums {
		total *= num // total = total * num
	}

	return total
}

func main() {

	fmt.Println("this", "is", "a", "variadic function")

	// call multiply

	fmt.Println(multiply(1, 2, 3, 4))

	// can be applied to slices

	nums := []int{1, 2, 3, 4}

	fmt.Println(multiply(nums...)) // don't forget ...

}
