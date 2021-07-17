package sliceutil

import (
	"fmt"
	"testing"
)

func TestInArray(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	if InArray(1, arr) {
		t.Log("YES")
	}
}

func TestArrayUnique(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5, 6}
	tarr := ArrayUnique(arr)
	fmt.Println(tarr)

}
