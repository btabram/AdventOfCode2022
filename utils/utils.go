package utils

import (
	"regexp"
)

// We don't expect any errors, treat them all as fatal.
func CheckErr[T any](val T, err error) T {
	if err != nil {
		panic(err)
	}
	return val
}

// There's no built-in set type but we can use a map with bool values and always insert true.
func Intersection[T comparable](a, b map[T]bool) map[T]bool {
	intersection := make(map[T]bool)
	for val := range a {
		if b[val] {
			intersection[val] = true
		}
	}
	return intersection
}

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
