package algorithms

import (
	"fmt"
	"testing"
)

func TestSearchMid(t *testing.T) {
	ar := []int{1, 2, 3}
	actual_value, err := BinarySearch(ar, 2)
	if err != nil && actual_value != 1 {
		t.Errorf("Search Failed")
	}

	ar_1 := []int{1}
	actual_value, err = BinarySearch(ar_1, 1)
	fmt.Println(actual_value, err)
	if err != nil && actual_value != 0 {
		t.Errorf("Search Failed")
	}

	ar_2 := []int{1, 2, 3, 4}
	actual_value, err = BinarySearch(ar_2, 3)
	if err != nil && actual_value != 2 {
		t.Errorf("Search Failed")
	}

}
func TestSearchLow(t *testing.T) {
	ar := []int{1, 2, 3}
	actual_value, err := BinarySearch(ar, 1)
	if err != nil {
		t.Errorf("Search Failed")
	}
	if actual_value != 0 {
		t.Errorf("Search Failed")
	}

	ar_2 := []int{1, 2, 3, 4}
	actual_value, err = BinarySearch(ar_2, 2)
	if err != nil && actual_value != 1 {
		t.Errorf("Search Failed")
	}
}
func TestSearchHigh(t *testing.T) {
	ar := []int{1, 2, 3}
	actual_value, err := BinarySearch(ar, 3)
	if err != nil {
		t.Errorf("Search Failed")
	}
	if actual_value != 2 {
		t.Errorf("Search Failed")
	}

	ar_2 := []int{1, 2, 3, 4}
	actual_value, err = BinarySearch(ar_2, 4)
	fmt.Println(actual_value, err)
	if err != nil && actual_value != 3 {
		t.Errorf("Search Failed")
	}

}

func TestSearchNotFoundMid(t *testing.T) {
	ar := []int{1, 2, 6}
	actual_value, err := BinarySearch(ar, 4)
	if err == nil {
		t.Errorf("Search Failed")
	}
	if actual_value != -1 {
		t.Errorf("Search Failed issue")
	}
}

func TestSearchNotFoundLow(t *testing.T) {
	ar := []int{1, 2, 3}
	actual_value, err := BinarySearch(ar, -1)
	if err == nil {
		t.Errorf("Search Failed")
	}
	if actual_value != -1 {
		t.Errorf("Search Failed issue")
	}
}

func TestSearchNotFoundHigh(t *testing.T) {
	ar := []int{1, 2, 3}
	actual_value, err := BinarySearch(ar, 4)
	if err == nil {
		t.Errorf("Search Failed")
	}
	if actual_value != -1 {
		t.Errorf("Search Failed issue")
	}
}
