package dstr_test

import (
	"DSA/dstr"
	"testing"
)

func TestArray(t *testing.T) {
	var arr dstr.Array[int]

	arr = dstr.AppendToArray(arr, 1)
	arr = dstr.AppendToArray(arr, 2)
	arr = dstr.AppendToArray(arr, 3)

	if dstr.GetFromArray(arr, 0) != 1 || dstr.GetFromArray(arr, 1) != 2 || dstr.GetFromArray(arr, 2) != 3 {
		t.Fatalf("There is an error witht the append function!")
	}

	if dstr.SizeOfArray(arr) != 3 {
		t.Fatalf("There is an error with the size function!")
	}

	dstr.SwapInArray(&arr, 0, 1)

	if dstr.GetFromArray(arr, 0) != 2 || dstr.GetFromArray(arr, 1) != 1 {
		t.Fatalf("There is an error witht the swap function!")
	}
}
