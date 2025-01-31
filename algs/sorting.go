package algs

import (
	"DSA/dstr"
	"fmt"
)

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func BubbleSort[T Ordered](arr dstr.Array[T]) dstr.Array[T] {
	size := dstr.Size(arr)

	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			val1 := dstr.Get[T](arr, j)
			val2 := dstr.Get[T](arr, j+1)
			if val1 > val2 {
				dstr.Swap[T](&arr, j, j+1)
			}
		}
	}

	return arr
}

func RevisedBubbleSort[T Ordered](arr dstr.Array[T]) dstr.Array[T] {
	size := dstr.Size(arr)

	for i := 0; i < size; i++ {
		isSwapped := false
		for j := 0; j < size-i-1; j++ {
			val1 := dstr.Get(arr, j)
			val2 := dstr.Get(arr, j+1)
			if val1 > val2 {
				dstr.Swap(&arr, j, j+1)
				isSwapped = true
			} else if !isSwapped {
				break
			}
		}
	}

	return arr
}

func SelectionSort[T comparable](arr dstr.Array[T]) dstr.Array[T] {
	size := dstr.Size(arr)
	for i := 0; i < size; i++ {
		fmt.Println(size)
	}
	return arr
}

func InsertionSort[T comparable](arr dstr.Array[T]) dstr.Array[T] {
	return arr
}

func QuickSort[T comparable](arr dstr.Array[T]) dstr.Array[T] {
	return arr
}
