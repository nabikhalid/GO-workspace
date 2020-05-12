package main

import "fmt"

func main() {

	i := 4

	switch i { // based on the value of this variable
	case 1: // if the value == 1
		fmt.Println("one")
	case 2: // if the vaue == 2
		fmt.Println("two")
	case 3: // if the value == 3
		fmt.Println("three")
	}

	j := 2

	switch j {
	case 1, 2: // j == 1 || j == 2
		fmt.Println("one or two")
	case 3:
		fmt.Println("three")
	default: // else
		fmt.Println("Not one, two, or three")
	}

	switch {

	case j == 5:
		fmt.Println(j, "is equal to 5")
	case j < 5:
		fmt.Println(j, "is less than 5")
	case j > 5:
		fmt.Println(j, "is greater than 5")

	}

}
