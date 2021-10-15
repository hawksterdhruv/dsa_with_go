package algorithms

import "fmt"

var count int

func BinarySearch(ar []int, target int) (int, error) {
	n := len(ar)
	mid := n / 2
	result := search(ar, target, mid, 0, n-1)

	if result == -1 {
		return -1, fmt.Errorf("result not found")
	} else {
		return result, nil
	}
}

func search(ar []int, target int, mid int, low int, high int) int {
	count++

	if count > 30 {
		return -1
	}

	fmt.Printf("%v,[%v],%v,%v,%v\n", ar, target, low, mid, high)

	if ar[mid] == target {
		fmt.Println("FOUND")
		return mid
	} else if high == low {
		return -1
	} else if ar[mid] > target {
		high = mid

	} else {
		low = mid + 1

	}
	mid = (low + high) / 2
	return search(ar, target, mid, low, high)
}
