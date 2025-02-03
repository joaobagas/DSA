package algs

import (
	"DSA/dstr"
)

// Interface to ensure that the variabls are numbers.
type ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func BubbleSort[T ordered](arr dstr.Array[T]) dstr.Array[T] {
	size := dstr.SizeOfArray(arr)

	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			val1 := dstr.GetFromArray(arr, j)
			val2 := dstr.GetFromArray(arr, j+1)
			if val1 > val2 {
				dstr.SwapInArray(&arr, j, j+1)
			}
		}
	}

	return arr
}

func RevisedBubbleSort[T ordered](arr dstr.Array[T]) dstr.Array[T] {
	size := dstr.SizeOfArray(arr)

	for i := 0; i < size; i++ {
		isSwapped := false
		for j := 0; j < size-i-1; j++ {
			val1 := dstr.GetFromArray(arr, j)
			val2 := dstr.GetFromArray(arr, j+1)
			if val1 > val2 {
				dstr.SwapInArray(&arr, j, j+1)
				isSwapped = true
			} else if !isSwapped {
				break
			}
		}
	}

	return arr
}

func SelectionSort[T ordered](arr dstr.Array[T]) dstr.Array[T] {
	size := dstr.SizeOfArray(arr)
	for i := 0; i < size; i++ {
		minIndex := i
		for j := i + 1; j < size; j++ {
			if dstr.GetFromArray(arr, i) > dstr.GetFromArray(arr, j) {
				minIndex = j
			}
		}
		dstr.SwapInArray(&arr, i, minIndex)
	}
	return arr
}

func InsertionSort[T ordered](arr dstr.Array[T]) dstr.Array[T] {
	size := dstr.SizeOfArray(arr)
	for i := 1; i < size; i++ {
		val := dstr.GetFromArray(arr, i)

		for j := i - 1; j >= 0; j-- {
			dstr.PrintArray(arr)
			if val < dstr.GetFromArray(arr, j) {
				arr = dstr.RemoveFromArray(arr, i)
				arr = dstr.InsertInArray(arr, j, val)
				break
			}
		}
	}

	return arr
}
