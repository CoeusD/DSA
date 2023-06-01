// Step 1 - Examples - 10 -> Not Prime, 5 -> Prime
// Step 2 - Utilitites - countFactors, check for prime
// Step 3 - Ranges
// Step 4 - Bruteforce
// Step 5 - Observation
// Step 6 - Edge Cases
package main

func main() {
	println(solve(5) == 1)
	println(solve(10) == 0)
}

func solve(A int64) int {
	numberOfFactors := countFactors(A)
	if isPrime(numberOfFactors) {
		return 1
	} else {
		return 0
	}
}

func countFactors(A int64) int {
	numberOfFactors := 0

	var i int64
	for i = 1; i*i <= A; i++ {
		if isFactor(A, i) {
			if i*i == A {
				numberOfFactors++
			} else {
				numberOfFactors += 2
			}
		}
	}
	return numberOfFactors
}

func isFactor(A int64, divisor int64) bool {
	return A%divisor == 0
}

func isPrime(numberOfFactors int) bool {
	return numberOfFactors == 2
}
