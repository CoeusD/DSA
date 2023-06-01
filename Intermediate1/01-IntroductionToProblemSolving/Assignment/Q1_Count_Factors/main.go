package main

func main() {
	println(solve(5) == 2)
	println(solve(10) == 4)
}

func solve(A int) int {
	numberOfFactors := 0
	for i := 1; i*i <= A; i++ {
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
