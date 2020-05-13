package main

import "fmt"

type rect struct {
	width, height int
}

func (r rect) area() int { // value type reciever
	return r.width * r.height
}

func (r *rect) perimeter() int { // pointaa
	return 2*r.width + 2*r.height
}

func main() {

	rectangle := rect{width: 10, height: 5}
	fmt.Println("area:", rectangle.area())

	rectanglePointer := &rectangle
	fmt.Println("perimeter:", rectanglePointer.perimeter())

}
