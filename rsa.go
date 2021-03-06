package rsa

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

// GenerateKeys generate the public and private keys for RSA-16 bit
// The function takes a seed as input and generate Public Key N, exponent e, Private key pk
func GenerateKeys(seed int) (int, int, int) {
	p, q := generateRandomPrimes(seed)
	N := p * q
	fmt.Println(p, q)
	phiN := (p - 1) * (q - 1)

	e := generateExponent(phiN, 1<<15)
	pk := modInverse(e, phiN)
	return N, e, pk
}

// generateExponent takes the phiN and a minimum value as input and after that
// increments till gcd(phiN, min) == 1.  While there might be a possible way to
// predit phiN via opposite calculation given that e is a public value and min
// is constant in our algorithm, I didn't consider this.
func generateExponent(phiN int, min int) int {
	for gcd(phiN, min) != 1 {
		min++
	}
	return min
}

// cgcd computes gsd of two given integers
func gcd(a int, b int) int {
	if a < b {
		a, b = b, a
	}
	for b > 1 {
		a, b = b, a%b
	}
	if b == 1 {
		return 1
	}
	return a
}

// modInverse takes to values (num and mod in sequence) and calculate the modulo inverse of num with mod
func modInverse(num int, mod int) int {
	num %= mod
	inverse := recurseModInverse(mod, num, 0, 1)
	for inverse < 0 {
		inverse += mod
	}
	return inverse % mod
}

func recurseModInverse(num1 int, num2 int, p0 int, p1 int) int {
	if num2 == 0 {
		return 0
	} else if num2 == 1 {
		return p1
	}

	p0, p1 = p1, (p0 - p1*(num1/num2))
	num1, num2 = num2, num1%num2
	return recurseModInverse(num1, num2, p0, p1)
}

func isPrime(num int) bool {
	for i := 2; i < int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func generateRandomPrimes(seed int) (int, int) {
	rand.Seed(int64(seed))
	min, max := 1<<15, 1<<16
	var p, q int
	for {
		p = rand.Intn(max-min) + min
		if isPrime(p) {
			break
		}
	}
	for !isPrime(q) {
		q = rand.Intn(max-min) + min
		if isPrime(q) {
			break
		}
	}

	return p, q
}

// SquareAndMultiply uses the well-known square and multiply algorithm to calculate the modulo exponent of a given number
// The function takes the inputs base, exponent and modulo in corresponding sequence
func SquareAndMultiply(base int, exp int, modulo int) int {
	res := base
	// converst the number to a string of 0 and 1 in binary
	bin := strconv.FormatInt(int64(exp), 2)
	for e := 1; e < len(bin); e++ {
		res *= res
		res %= modulo
		if bin[e] == '1' {
			res *= base
			res %= modulo
		}
	}
	return res
}
