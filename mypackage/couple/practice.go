package couple

import "fmt"

func Couple() {
	var start, end int

	fmt.Print("عدد اول رو وارد بکنید")
	fmt.Scan(&start)
	fmt.Print("عدد پایان رو وارد کنید")
	fmt.Scan(&end)
	eventnumber := []int{}
	for i := start; i <= end; i++ {
		if i%2 == 0 {
			eventnumber = append(eventnumber, i)
		}
	}
	fmt.Println(eventnumber)
}
