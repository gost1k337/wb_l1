// Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {
	a := 5
	b := 10

	a, b = b, a

	fmt.Println(a)
	fmt.Println(b)
}
