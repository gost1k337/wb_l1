package main

import (
	"fmt"
	"math"
	"math/big"
)

func calculate(firstValue, secondValue *big.Float, operator string) *big.Float {
	result := new(big.Float)

	switch operator {

	case "+":
		return result.Add(firstValue, secondValue)

	case "-":
		return result.Sub(firstValue, secondValue)

	case "*":
		return result.Mul(firstValue, secondValue)

	case "/":
		return result.Quo(firstValue, secondValue)

	default:
		return result
	}
}

func main() {
	a, b := big.NewFloat(math.Pow(2, 21)), big.NewFloat(math.Pow(2, 25)) // инициализируем значения float >2^20, стандартная точность - 53, режим округления - до ближайшего целого

	fmt.Println(calculate(a, b, "+"))
	fmt.Println(calculate(a, b, "-"))
	fmt.Println(calculate(a, b, "*"))
	fmt.Println(calculate(a, b, "/"))

}
