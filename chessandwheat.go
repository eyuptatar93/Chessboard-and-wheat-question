package main

import "fmt"

func main() {
	num := []uint64{}
	var total uint64 = 0

	fmt.Println(num)
	for i := 1; i <= 64; i++ {
		if i == 1 {
			num = append(num, uint64(i))
		} else {
			num = append(num, 2*(num[len(num)-1]))
			// get last item
		}
	}
	for _, s := range num {
		fmt.Println(s)
	}
	for _, s := range num {
		total += s
	}
	fmt.Println("Total: ", total)
}
