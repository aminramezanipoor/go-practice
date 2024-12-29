package main

import (
	"couples/mypackage/couple"
	"fmt"
)

func main() {
	var choice int
	couple.Couple()
	for {
		fmt.Println("\n--- student management system ---")
		fmt.Println("1. add student")
		fmt.Println("2. exit")
		fmt.Print("enter your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			var name string
			var score int
			fmt.Print("enter student name :")
			fmt.Scan(&name)
			fmt.Print("enter student score :")
			fmt.Scan(&score)

			if score >= 50 {
				fmt.Printf("studen %s has passed.\n", name)
			} else {
				fmt.Printf("studen %s has failed.\n", name)
			}
		case 2:
			fmt.Println("exiting the program...")
			return
		default:
			fmt.Println("invalid choice")
		}
	}

}
