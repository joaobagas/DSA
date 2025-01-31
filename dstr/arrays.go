package dstr

import (
	"fmt"
)

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
