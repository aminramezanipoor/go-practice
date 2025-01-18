package main

import "fmt"

func main() {
	number := 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			number += i
		}
	}
	fmt.Println(number)
}
