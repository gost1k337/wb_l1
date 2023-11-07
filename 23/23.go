package main

import (
	"fmt"
	"reflect"
)

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	valid := []int{1, 2, 3, 5, 6, 7, 8}

	slice = remove(slice, 3)

	fmt.Println(reflect.DeepEqual(slice, valid))
}
