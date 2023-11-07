// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
//Приведите корректный пример реализации.
//
//var justString string
//
//func someFunc() {
//	v := createHugeString(1 << 10)
//	justString = v[:100] - выход за пределы слайса если длина слайса меньше 100
//}

package main

import (
	"fmt"
	"math/rand"
)

var justString string

func someFunc() {
	v := createHugeString(99)
	if len(v) < 100 {
		justString = v
	} else {
		justString = v[:100]
	}
}

func createHugeString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randInt(97, 122))
	}

	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	someFunc()
	fmt.Println(justString)
}
