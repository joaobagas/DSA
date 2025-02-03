package dstr

import (
	"fmt"
)

// A linear data structure that stores a collection of elements of the same data type.
type Array[T any] struct {
	sli []T
}

// The AppendToArray function adds elements to the array.
func AppendToArray[T any](arr Array[T], el T) Array[T] {
	arr.sli = append(arr.sli, el)
	return arr
}

// Removes an element from an array.
func RemoveFromArray[T any](arr Array[T], index int) Array[T] {
	size := len(arr.sli)
	newSli := make([]T, size-1)

	// Copy elements before the index
	copy(newSli, arr.sli[:index])

	// Copy elements after the index
	copy(newSli[index:], arr.sli[index+1:])

	arr.sli = newSli
	return arr
}

// Returns the element of the array.
func GetFromArray[T any](arr Array[T], index int) T {
	return arr.sli[index]
}

// Gives the size of the array.
func SizeOfArray[T any](arr Array[T]) int {
	return len(arr.sli)
}

// Prints the information inside the array.
func PrintArray[T any](arr Array[T]) {
	fmt.Println(arr.sli)
}

// Swaps two elements of an array.
func SwapInArray[T any](arr *Array[T], firstIndex int, secondIndex int) {
	temp := arr.sli[firstIndex]
	arr.sli[firstIndex] = arr.sli[secondIndex]
	arr.sli[secondIndex] = temp
}

func InsertInArray[T any](arr Array[T], index int, val T) Array[T] {
	size := len(arr.sli)
	newSli := make([]T, size+1)

	// Copy before index
	copy(newSli, arr.sli[:index])

	newSli[index] = val

	// Copy after index
	copy(newSli[index+1:], arr.sli[index:])

	arr.sli = newSli
	return arr
}
