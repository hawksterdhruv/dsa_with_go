package algorithms

import "fmt"

func InsertionSort(ar []int) []int {
	n := len(ar)
	for i := 1; i < n; i++ {
		fmt.Println(ar)
		for j := i - 1; j >= 0; j-- {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
			}
		}
	}
	return ar
}
