package main

import (
	"DSA/algs"
	"DSA/dstr"
)

func main() {
	var arr dstr.Array[int]

	arr = dstr.AppendToArray(arr, 6)
	arr = dstr.AppendToArray(arr, 4)
	arr = dstr.AppendToArray(arr, 3)

	arr = algs.InsertionSort(arr)
}
