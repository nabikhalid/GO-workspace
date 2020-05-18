package main

import (
	"fmt"

	"github.com/nabikhalid/GO-workspace/packages/mymath"
)

// all this time the answer was -u
// go get -u "github.com/..."

func main() {

	fmt.Println("yessirski")
	fmt.Println(mymath.Add(1, 2))

	// lowercase doesn't work
	// fmt.Println(mymath.subtract(1, 2))
}
