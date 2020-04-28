package main

import "fmt"

func main() {

	ballon := make(map[string]int)

	ballon["Ronaldo"] = 5
	ballon["Messi"] = 6

	fmt.Println(ballon)

	fmt.Println(ballon["Messi"])

	delete(ballon, "Ronaldo")

	fmt.Println(ballon)

}
