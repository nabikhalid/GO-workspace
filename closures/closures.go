package main

import "fmt"

func sayHello(msg string) {
	fmt.Println(msg)
}

func returnAnon() func(string) {
	// return an anonymous function
	return func(msg string) {
		fmt.Println(msg)
	}
}
func main() {
	sayHello("hello")

	// anonymous function
	func(msg string) {
		fmt.Println(msg)
	}("Hello Anonymous")

	printFunction := returnAnon()
	printFunction("Hello returned from anonymous")

	// end of anonymous functions, definatley need to review this

}
