package main

import "fmt"

func main() {
	total := uint64(0)
	arr := []uint64{}
	
	for i := 0; i < 64; i++ {
		if i == 0 {
			total += 1
			arr = append(arr, 1)
		} else {
			total += 2 * arr[len(arr)-1]
			arr = append(arr, (2 * arr[len(arr)-1]))
		}
	}

	fmt.Println(total)
	fmt.Println(arr)
}
