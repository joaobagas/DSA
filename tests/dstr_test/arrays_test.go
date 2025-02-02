package dstr_test

import (
	"DSA/dstr"
	"testing"
)

func TestArray(t *testing.T) {
	var arr dstr.Array[int]

	arr = dstr.Append(arr, 1)
	arr = dstr.Append(arr, 2)
	arr = dstr.Append(arr, 3)

	if dstr.Get(arr, 0) != 1 || dstr.Get(arr, 1) != 2 || dstr.Get(arr, 2) != 3 {
		t.Fatalf("There is an error witht the append function!")
	}

	if dstr.Size(arr) != 3 {
		t.Fatalf("There is an error with the size function!")
	}

	dstr.Swap(&arr, 0, 1)

	if dstr.Get(arr, 0) != 2 || dstr.Get(arr, 1) != 1 {
		t.Fatalf("There is an error witht the swap function!")
	}
}
