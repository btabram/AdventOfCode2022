package utils

import "regexp"

func Lines(str string) []string {
	// Use a regex to make this work for both Windows and Unix line endings.
	return regexp.MustCompile("\r?\n").Split(str, -1)
}

func Sum(slice []int) int {
	sum := 0
	for _, item := range slice {
		sum += item
	}
	return sum
}

func Transform[A, B any](slice []A, transformFn func(A) B) []B {
	newSlice := make([]B, len(slice))
	for i := range slice {
		newSlice[i] = transformFn(slice[i])
	}
	return newSlice
}
