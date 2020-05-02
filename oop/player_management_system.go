package main

import (
	"fmt"
)

type player struct {
	ID string
	ratio float32
}

func getID (p player) string{
	return p.ID
}

func setID (p *player, newID string) { // accept pointer *
	p.ID = newID
}

func getratio (p player) float32{
	return p.ratio
}

func setratio (p *player, newratio float32) {
	p.ratio = newratio
}

func main() {

	// p1 := player{ID: "NK", ratio: 1.50}
	// // players := []player{}

	// fmt.Println(getID(p1))
	// setID(&p1, "KSav")

	// fmt.Println(getID(p1))

	fmt.Println("Welcome to your Call of Duty: Modern Warfare database.")
	fmt.Println("What would you like to do?")
	
	var response string 

	fmt.Scanln(&response)

	for response != "exit" {
		
		if response == "add player" { 
			

			fmt.Println("What would you like to do?")
			fmt.Scanln(&response)
		} else if response == "remove player" { 


			fmt.Println("What would you like to do?")
			fmt.Scanln(&response)
		} else if response == "" {


			fmt.Println("What would you like to do?")
			fmt.Scanln(&response)
		} else if response == "" {


			fmt.Println("What would you like to do?")
			fmt.Scanln(&response)
		} else if response == "" {


			fmt.Println("What would you like to do?")
			fmt.Scanln(&response)
		} else if response == "" {


			fmt.Println("What would you like to do?")
			fmt.Scanln(&response)
		} else { 
			fmt.Println("What would you like to do?")
			fmt.Scanln(&response)
		}
		
	}
	
	fmt.Println("Have a good day!")

}