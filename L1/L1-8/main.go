package main

import "fmt"

func setBitToOne(number int64, bitPosition uint) int64 {
	mask := int64(1) << bitPosition
	return number | mask
}

func setBitToZero(number int64, bitPosition uint) int64 {
	mask := int64(1) << bitPosition
	return number &^ mask
}

func getBit(number int64, bitPosition uint) int64 {
	return (number >> bitPosition) & 1
}

func printBinary(number int64, bits int) {
	for i := bits - 1; i >= 0; i-- {
		fmt.Printf("%d", (number>>i)&1)
		if i%4 == 0 && i != 0 {
			fmt.Print(" ")
		}
	}
}

func main() {
	var number int64 = 5
	
	fmt.Println("=== Initial number ===")
	fmt.Printf("Decimal: %d\n", number)
	fmt.Print("Binary:  ")
	printBinary(number, 8)
	fmt.Println()

	fmt.Println("=== Set bit 1 to 0 ===")
	resultZero := setBitToZero(number, 1)
	fmt.Printf("Decimal: %d\n", resultZero)
	fmt.Print("Binary:  ")
	printBinary(resultZero, 8)
	fmt.Println()

	fmt.Println("=== Set bit 1 to 1 ===")
	resultOne := setBitToOne(number, 1)
	fmt.Printf("Decimal: %d\n", resultOne)
	fmt.Print("Binary:  ")
	printBinary(resultOne, 8)
	fmt.Println()

	fmt.Println("=== More examples ===")
	testNumber := int64(42)
	fmt.Printf("Original: %d (binary: ", testNumber)
	printBinary(testNumber, 8)
	fmt.Println(")")

	fmt.Printf("Set bit 0 to 1: %d (binary: ", setBitToOne(testNumber, 0))
	printBinary(setBitToOne(testNumber, 0), 8)
	fmt.Println(")")

	fmt.Printf("Set bit 3 to 0: %d (binary: ", setBitToZero(testNumber, 3))
	printBinary(setBitToZero(testNumber, 3), 8)
	fmt.Println(")")

	fmt.Printf("Set bit 5 to 1: %d (binary: ", setBitToOne(testNumber, 5))
	printBinary(setBitToOne(testNumber, 5), 8)
	fmt.Println(")")

	fmt.Println("=== Checking individual bits ===")
	checkNumber := int64(85)
	fmt.Printf("Number: %d (binary: ", checkNumber)
	printBinary(checkNumber, 8)
	fmt.Println(")")
	
	for bitPos := uint(0); bitPos < 8; bitPos++ {
		bitValue := getBit(checkNumber, bitPos)
		fmt.Printf("Bit %d: %d\n", bitPos, bitValue)
	}
}
