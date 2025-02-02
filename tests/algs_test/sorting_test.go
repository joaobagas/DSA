package algs_test

import (
	"DSA/algs"
	"DSA/dstr"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var arr dstr.Array[int]

	arr = dstr.Append(arr, 6)
	arr = dstr.Append(arr, 4)
	arr = dstr.Append(arr, 3)

	arr = algs.BubbleSort(arr)

	if dstr.Get(arr, 0) != 3 || dstr.Get(arr, 1) != 4 || dstr.Get(arr, 2) != 6 {
		t.Fatalf("There is an error witht the bubble sort function!")
	}
}
