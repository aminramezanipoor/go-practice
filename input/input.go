package input_1

import "fmt"

func Input() {
	var name string
	var age int

	fmt.Scan(&name, &age)
	fmt.Printf("name: %s\n age: %d\n", name, age)
}
