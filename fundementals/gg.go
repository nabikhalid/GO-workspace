package main

import ("fmt")

// go run file_name.go

func main()  {
	
	fmt.Println("what's up")

	var x int = 5
	fmt.Println(x);

	var y int = 7
	fmt.Println(y);

	var z int
	fmt.Println(z); // zero value is 0

	var sum int = x + y

	fmt.Println(sum)

	infer := 7 // infers type
	fmt.Println(infer); 

	if infer > 6 { // no brackets
		fmt.Println("More than six")
	} else if x == 5 { // EXACT bracket format
		fmt.Println("yessssir")
	} else { // EXACT bracket format
		fmt.Println("nosir")
	}

}