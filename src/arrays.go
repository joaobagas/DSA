package src

import (
	"fmt"
)

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

type Array[T any] struct {
	sli []T
}

func Append[T any](arr Array[T], el T) Array[T] {
	arr.sli = append(arr.sli, el)
	return arr
}

func Remove[T any](arr Array[T], index int) Array[T] {
	newSli := make([]T, len(arr.sli)-1)
	for i := range arr.sli {
		if i != index {
			newSli = append(newSli, arr.sli[i])
		}
	}
	arr.sli = newSli
	return arr
}

func Get[T any](arr Array[T], index int) T {
	return arr.sli[index]
}

func Size[T any](arr Array[T]) int {
	return len(arr.sli)
}

func Print[T any](arr Array[T]) {
	fmt.Println(arr.sli)
}

func Swap[T any](arr *Array[T], firstIndex int, secondIndex int) {
	temp := arr.sli[firstIndex]
	arr.sli[firstIndex] = arr.sli[secondIndex]
	arr.sli[secondIndex] = temp
}

func BubbleSort[T Ordered](arr Array[T]) Array[T] {
	size := Size(arr)

	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			val1 := Get[T](arr, j)
			val2 := Get[T](arr, j+1)
			if val1 > val2 {
				Swap[T](&arr, j, j+1)
			}
		}
	}

	return arr
}

func RevisedBubbleSort[T Ordered](arr Array[T]) Array[T] {
	size := Size(arr)

	for i := 0; i < size; i++ {
		isSwapped := false
		for j := 0; j < size-i-1; j++ {
			val1 := Get(arr, j)
			val2 := Get(arr, j+1)
			if val1 > val2 {
				Swap(&arr, j, j+1)
				isSwapped = true
			} else if !isSwapped {
				break
			}
		}
	}

	return arr
}

func SelectionSort[T comparable](arr Array[T]) Array[T] {
	size := Size(arr)
	for i := 0; i < size; i++ {
		fmt.Println(size)
	}
	return arr
}

func InsertionSort[T comparable](arr Array[T]) Array[T] {
	return arr
}

func QuickSort[T comparable](arr Array[T]) Array[T] {
	return arr
}
