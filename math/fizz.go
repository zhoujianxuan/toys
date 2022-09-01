package math

import "strconv"

func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := range res {
		v := i + 1
		if v%15 == 0 {
			res[i] = "FizzBuzz"
		} else if v%3 == 0 {
			res[i] = "Fizz"
		} else if v%5 == 0 {
			res[i] = "Buzz"
		} else {
			res[i] = strconv.Itoa(v)
		}
	}
	return res
}
