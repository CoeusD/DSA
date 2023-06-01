// Step 1 - Example - 4 -> 2, 1001 -> -1
// Step 2 - Utilities - None
// Step 3 - Ranges - 1 to i * i <= A
// Step 4 - Bruteforce

package main

func main() {
	println(solve(4) == 2)
	println(solve(1001) == -1)
}

func solve(A int) int {
	for i := 1; i*i <= A; i++ {
		if i*i == A {
			return i
		}
	}
	return -1
}
