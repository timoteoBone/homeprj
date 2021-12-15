package fizzbuzz

import "strconv"

var fizzRes string = "fizz"
var buzzRes string = "buzz"
var FizzBuzzRes string = "FizzBuzz"

func FizzBuzz(num int) string {
	var res string
	fizz := dividendBy3(num)
	buzz := dividendBy5(num)

	if fizz && buzz {
		res = FizzBuzzRes
	} else if fizz && !buzz {
		res = fizzRes
	} else if buzz && !fizz {
		res = buzzRes
	} else {
		res = strconv.Itoa(num)
	}
	return res
}

func dividendBy3(num int) bool {
	return num%3 == 0

}

func dividendBy5(num int) bool {
	return num%5 == 0
}
