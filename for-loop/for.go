package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	SecretNumber := rand.Intn(10) + 1
	fmt.Println("یک عدد بین 1 تا 10 انتخاب کن")
	var guess int
	for {
		fmt.Print("حدس شما:")
		fmt.Scan(&guess)

		if guess > SecretNumber {
			fmt.Println("عدد کوچکتر است")
		} else if guess < SecretNumber {
			fmt.Println("عدد بزرگتر است")
		} else {
			fmt.Println("حدس شما درست است")
			break
		}
	}
}
