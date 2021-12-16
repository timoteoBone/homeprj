package calculus

import (
	errors "github.com/timoteoBone/homeprj/newErrors"
)

func Sumar(nums []int) int {
	suma := 0
	for _, num := range nums {
		suma += num
	}
	return suma
}

func Restar(nums []int) int {
	resta := 0
	for _, num := range nums {
		resta -= num
	}
	return resta
}

func Dividir(num1, num2 int) (int, error) {
	if num2 != 0 {
		return num1 / num2, nil
	}
	return 0, &errors.InternalError{}
}

func Multiplicar(nums []int) int {
	multiplicacion := nums[0]
	for _, num := range nums[1:] {
		multiplicacion *= num
	}
	return multiplicacion
}
