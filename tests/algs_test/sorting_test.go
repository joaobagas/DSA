package algs_test

import (
	"DSA/algs"
	"DSA/dstr"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var arr dstr.Array[int]

	arr = dstr.AppendToArray(arr, 6)
	arr = dstr.AppendToArray(arr, 4)
	arr = dstr.AppendToArray(arr, 3)

	arr = algs.BubbleSort(arr)

	if dstr.GetFromArray(arr, 0) != 3 || dstr.GetFromArray(arr, 1) != 4 || dstr.GetFromArray(arr, 2) != 6 {
		t.Fatalf("There is an error witht the bubble sort function!")
	}
}

func TestBubbleSortRevised(t *testing.T) {
	var arr dstr.Array[int]

	arr = dstr.AppendToArray(arr, 6)
	arr = dstr.AppendToArray(arr, 4)
	arr = dstr.AppendToArray(arr, 3)

	arr = algs.RevisedBubbleSort(arr)

	if dstr.GetFromArray(arr, 0) != 3 || dstr.GetFromArray(arr, 1) != 4 || dstr.GetFromArray(arr, 2) != 6 {
		t.Fatalf("There is an error witht the revised bubble sort function!")
	}
}

func TestSelectionSort(t *testing.T) {
	var arr dstr.Array[int]

	arr = dstr.AppendToArray(arr, 6)
	arr = dstr.AppendToArray(arr, 4)
	arr = dstr.AppendToArray(arr, 3)

	arr = algs.SelectionSort(arr)

	if dstr.GetFromArray(arr, 0) != 3 || dstr.GetFromArray(arr, 1) != 4 || dstr.GetFromArray(arr, 2) != 6 {
		t.Fatalf("There is an error witht the selection sort function!")
	}
}

func TestInsertionSort(t *testing.T) {
	var arr dstr.Array[int]

	arr = dstr.AppendToArray(arr, 6)
	arr = dstr.AppendToArray(arr, 4)
	arr = dstr.AppendToArray(arr, 3)

	arr = algs.InsertionSort(arr)

	if dstr.GetFromArray(arr, 0) != 3 || dstr.GetFromArray(arr, 1) != 4 || dstr.GetFromArray(arr, 2) != 6 {
		t.Fatalf("There is an error witht the insertion sort function!")
	}
}
