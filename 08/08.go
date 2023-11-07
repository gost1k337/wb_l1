// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import "fmt"

func toOne(n int64, i int) int64 {
	n = n | (1 << (i - 1)) // сдвигаем бит на i - 1 позиций влево и применяем побитовое ИЛИ к n
	return n
}

func toZero(n int64, i int) int64 {
	n = n & ^(1 << (i - 1)) // сдвигаем бит на i - 1 позиций влево и применяем побитовое И к n
	return n
}

func main() {
	var n int64
	n = 8
	fmt.Println(toOne(n, 4))
	fmt.Println(toZero(n, 4))
}
