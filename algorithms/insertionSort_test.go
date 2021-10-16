package algorithms

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	ar := []int{5, 7, 2, 1, 8}
	fmt.Println("START Insertion")
	actual := InsertionSort(ar)
	expected := []int{1, 2, 5, 7, 8}
	fmt.Println(actual)
	fmt.Println("END")
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Sorting Failed")
	}

}
