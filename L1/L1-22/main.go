package main

import (
	"fmt"
	"math/big"
)

// способ 1: обёртки над big.Int
func addBigInt(a, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

func subBigInt(a, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func mulBigInt(a, b *big.Int) *big.Int {
	return new(big.Int).Mul(a, b)
}

func divBigInt(a, b *big.Int) *big.Int {
	return new(big.Int).Div(a, b)
}

// способ 2: строковая арифметика (сложение)
func addStrings(a, b *big.Int) *big.Int {
	strA, strB := a.String(), b.String()
	result := []byte{}
	carry := 0
	i, j := len(strA)-1, len(strB)-1

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(strA[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(strB[j] - '0')
			j--
		}
		carry = sum / 10
		result = append([]byte{byte(sum%10 + '0')}, result...)
	}

	res, _ := new(big.Int).SetString(string(result), 10)
	return res
}

// способ 3: срез цифр (умножение столбиком)
func mulDigits(a, b *big.Int) *big.Int {
	strA, strB := a.String(), b.String()
	if strA == "0" || strB == "0" {
		return big.NewInt(0)
	}

	result := make([]int, len(strA)+len(strB))

	for i := len(strA) - 1; i >= 0; i-- {
		for j := len(strB) - 1; j >= 0; j-- {
			mul := int(strA[i]-'0') * int(strB[j]-'0')
			pos := i + j + 1
			sum := mul + result[pos]
			result[pos] = sum % 10
			result[i+j] += sum / 10
		}
	}

	str := ""
	for _, v := range result {
		if !(len(str) == 0 && v == 0) {
			str += string(byte(v + '0'))
		}
	}

	res, _ := new(big.Int).SetString(str, 10)
	return res
}

func main() {
	a := big.NewInt(1 << 21)
	b := big.NewInt(1 << 20)

	fmt.Println("способ 1 (big.Int):")
	fmt.Println("add:", addBigInt(a, b))
	fmt.Println("sub:", subBigInt(a, b))
	fmt.Println("mul:", mulBigInt(a, b))
	fmt.Println("div:", divBigInt(a, b))

	fmt.Println("\nспособ 2 (строковое сложение):")
	fmt.Println("add:", addStrings(a, b))

	fmt.Println("\nспособ 3 (умножение столбиком):")
	fmt.Println("mul:", mulDigits(a, b))

	// очень большие числа
	x, _ := new(big.Int).SetString("99999999999999999999", 10)
	y, _ := new(big.Int).SetString("11111111111111111111", 10)

	fmt.Println("\nбольшие числа (все способы):")
	fmt.Println("big.Int add:", addBigInt(x, y))
	fmt.Println("string add:", addStrings(x, y))
	fmt.Println("digits mul:", mulDigits(x, y))
}
