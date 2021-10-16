package algorithms

import "fmt"

func SelectionSort(ar []int) []int {
	n := len(ar)
	for i := 0; i < n; i++ {
		fmt.Println(ar)
		for j := i + 1; j < n; j++ {
			if ar[i] > ar[j] {
				ar[i], ar[j] = ar[j], ar[i]
			}
		}
	}
	return ar
}
