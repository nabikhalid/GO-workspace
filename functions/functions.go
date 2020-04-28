package main 
import ( 
	"fmt" 
	"errors" 
	"math" 
)

func main() {

	result := sum(7, 7)
	fmt.Println(result)

	// test sqrt

}

func sum (x int, y int) int { // func name (input type) return type

	return x + y

}

func sqrt(x float64) (float64, error) { // functions can have multiple return types

	if x < 0 {
		return 0, errors.New("Undefined for negative numbers.")
	}

	return math.Sqrt(x), nil

}