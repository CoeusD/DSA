package main

func main() {
	println(solve(5) == 1)
	println(solve(10) == 0)
}

func solve(A int64) int {
	numberOfFactors := countNumberOfFactors(A)
	if numberOfFactors == 2 {
		return 1
	} else {
		return 0
	}
}

func countNumberOfFactors(A int64) int {
	numberOfFactors := 0

	var i int64
	for i = 1; i*i <= A; i++ {
		if A%i == 0 {
			if i*i == A {
				numberOfFactors++
			} else {
				numberOfFactors += 2
			}
		}
	}
	return numberOfFactors
}
