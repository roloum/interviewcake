package nthfibonacci

import "errors"

var errNegativeInt = errors.New("Number can't be negative")

func nthFibonacci(n int) (int, error) {

	if n < 0 {
		return 0, errNegativeInt
	} else if n < 2 { //fib(0)=0, fib(1)=1
		return n, nil
	}

	fib, last := 1, 0
	for i := 2; i <= n; i++ {
		last, fib = fib, fib+last
	}

	return fib, nil
}
