package main

import "fmt"

func printArray(arr *[]int) {
	fmt.Println("Current array", *arr)
}
func addNumber(arr *[]int, num int) {
	*arr = append(*arr, num)
	fmt.Printf("The number %d was added", num)
}

func removeNumber(arr *[]int, num int) {
	for i, valu := range *arr {
		if valu == num {
			*arr = append((*arr)[:i], (*arr)[i+1:]...)
			fmt.Printf("The number %d was removed \n", num)
			return
		}
	}
	fmt.Printf("The number %d was not removed \n", num)
}

func searchNumber(arr *[]int, num int) {
	for i, valu := range *arr {
		if valu == num {
			fmt.Printf("The number %d was found in position %d \n", num, i)
			return
		}
	}
	fmt.Printf("The number %d does not exist in the array\n", num)
}

func main() {
	arr := []int{}

	for {
		fmt.Println("\nChoose")
		fmt.Println("1. add number")
		fmt.Println("2. remove number")
		fmt.Println("3. Number search")
		fmt.Println("4. Array display")
		fmt.Println("5. exit")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter number")
			var num int
			fmt.Scan(&num)
			addNumber(&arr, num)
		case 2:
			fmt.Print("The number to be deleted")
			var num int
			fmt.Scan(&num)
			removeNumber(&arr, num)
		case 3:
			fmt.Print("The desired number to search for")
			var num int
			fmt.Scan(&num)
			searchNumber(&arr, num)
		case 4:
			printArray(&arr)
		case 5:
			fmt.Println("Goodbye")
			return
		default:
			fmt.Println("Invalid")
		}
	}
}
