package main

import "fmt"

func swapWithAddition(a, b int) (int, int) {
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

func swapWithSubtraction(a, b int) (int, int) {
	a = a - b
	b = a + b
	a = b - a
	return a, b
}

func swapWithXOR(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}

func swapWithMultiplication(a, b int) (int, int) {
	if a == 0 || b == 0 {
		return b, a
	}
	a = a * b
	b = a / b
	a = a / b
	return a, b
}

func main() {
	firstNumber := 10
	secondNumber := 20
	
	fmt.Println("=== Swap with addition/subtraction ===")
	fmt.Printf("Before: a = %d, b = %d\n", firstNumber, secondNumber)
	resultA, resultB := swapWithAddition(firstNumber, secondNumber)
	fmt.Printf("After: a = %d, b = %d\n", resultA, resultB)
	fmt.Println()
	
	fmt.Println("=== Swap with subtraction ===")
	fmt.Printf("Before: a = %d, b = %d\n", firstNumber, secondNumber)
	resultA, resultB = swapWithSubtraction(firstNumber, secondNumber)
	fmt.Printf("After: a = %d, b = %d\n", resultA, resultB)
	fmt.Println()
	
	fmt.Println("=== Swap with XOR ===")
	fmt.Printf("Before: a = %d, b = %d\n", firstNumber, secondNumber)
	resultA, resultB = swapWithXOR(firstNumber, secondNumber)
	fmt.Printf("After: a = %d, b = %d\n", resultA, resultB)
	fmt.Println()
	
	fmt.Println("=== Swap with multiplication/division ===")
	fmt.Printf("Before: a = %d, b = %d\n", firstNumber, secondNumber)
	resultA, resultB = swapWithMultiplication(firstNumber, secondNumber)
	fmt.Printf("After: a = %d, b = %d\n", resultA, resultB)
	fmt.Println()
	
	fmt.Println("=== Testing with negative numbers ===")
	negativeA := -15
	negativeB := 25
	fmt.Printf("Before: a = %d, b = %d\n", negativeA, negativeB)
	resultA, resultB = swapWithXOR(negativeA, negativeB)
	fmt.Printf("After (XOR): a = %d, b = %d\n", resultA, resultB)
	fmt.Println()
	
	fmt.Println("=== Testing with zero ===")
	zeroA := 0
	zeroB := 42
	fmt.Printf("Before: a = %d, b = %d\n", zeroA, zeroB)
	resultA, resultB = swapWithXOR(zeroA, zeroB)
	fmt.Printf("After (XOR): a = %d, b = %d\n", resultA, resultB)
	fmt.Println()
	
	fmt.Println("=== XOR step-by-step explanation ===")
	x := 5
	y := 9
	fmt.Printf("Initial: x = %d (%04b), y = %d (%04b)\n", x, x, y, y)
	x = x ^ y
	fmt.Printf("x = x ^ y: x = %d (%04b), y = %d (%04b)\n", x, x, y, y)
	y = x ^ y
	fmt.Printf("y = x ^ y: x = %d (%04b), y = %d (%04b)\n", x, x, y, y)
	x = x ^ y
	fmt.Printf("x = x ^ y: x = %d (%04b), y = %d (%04b)\n", x, x, y, y)
}
