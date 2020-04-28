package main

import "fmt"

func main() {

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// no while loops, instead ... 


	j := 0

	for j < 5 {
		fmt.Println(j)
		j++
	}

	// prints same thing 

	array := []string{"Catalanes", "Real Madrid", "Barcelona"}

	for index, value := range array { // loop through array
		fmt.Println("index:", index, "value:", value) // string concat 
	}

	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m { // loop through map
		fmt.Println("key:", key, "value:", value)
	}

}
