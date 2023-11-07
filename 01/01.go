// Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов
// в структуре Action от родительской структуры Human (аналог наследования).

package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) Walk() {
	fmt.Println("human: walk()")
}

// Встраиваем структуру
type Action struct {
	Human
}

// Вызываем метод "родителя"
func (a Action) Walk() {
	a.Human.Walk()
}

func main() {
	h := Human{"Alex", 23}
	a := Action{h}

	fmt.Println("Human:")
	h.Walk()

	fmt.Println("--------------")

	fmt.Println("Action:")
	a.Walk()
}
