// Step 1 - Example - 10 -> 1,2,5,10
// Step 2 - Utilities - IsFactor
// Step 3 - Ranges - 1 to N
// Step 4 - Bruteforce - For all i isFactor(A, i)
// Step 5 - Observation - Check till sqrt
// Step 6 - Edge Cases

package main

func main() {
	println(solve(5) == 2)
	println(solve(10) == 4)
}

func solve(A int) int {
	numberOfFactors := 0
	for i := 1; i*i <= A; i++ {
		if IsFactor(A, i) {
			if i*i == A {
				numberOfFactors++
			} else {
				numberOfFactors += 2
			}
		}
	}
	return numberOfFactors
}

func IsFactor(A int, divisor int) bool {
	return A%divisor == 0
}
