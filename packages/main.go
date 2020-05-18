package main

import (
	"fmt"

	"github.com/nabikhalid/GO-workspace/packages/mymath"
)

// HOW TO IMPORT PACKAGES IN GO:
// go get -u "github.com/..." in the folder containing main, wow.

func main() {

	fmt.Println("yessirski")
	fmt.Println(mymath.Add(1, 2))
	fmt.Println(mymath.Yessirski())

	// lowercase doesn't work
	// fmt.Println(mymath.subtract(1, 2))
}
