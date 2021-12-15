package main

import (
	"fmt"
	"os"
	"strconv"

	calculate "github.com/timoteoBone/homeprj/8-calculator"
)

func main() {
	// write firt the operator we want to use, then all the numbers we want
	nums := []int{}
	for _, num := range os.Args[2:] {
		aux, err := strconv.Atoi(num)
		if err == nil {
			nums = append(nums, aux)
		}
	}

	result, err := operate(os.Args[1], nums)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func operate(operador string, numeros []int) (int, error) {
	var error error
	var result int
	switch operador {
	case "+":
		result = calculate.Sumar(numeros)
	case "-":
		result = calculate.Restar(numeros)
	case "/":
		result, error = calculate.Dividir(numeros[0], numeros[1])
	case "++":
		result = calculate.Multiplicar(numeros)
	}

	return result, error
}
