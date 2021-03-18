package rsa

// GenerateKeys generate the public and private keys for RSA-16 bit
func GenerateKeys(int p, int q) (int, int, int) {
	// Very random, believe me
	// p, q := 34667, 44207
	N := p * q
	phiN := (p - 1) * (q - 1)

	// Very bad practice, but whatever
	e := phiN - 1
	pk := modInverse(e, phiN)
	return N, e, pk
}

// modInverse takes to values (num and mod in sequence) and calculate the modulo inverse of num with mod 
func modInverse(num int, mod int) int {
	num %= mod
	inverse := recurseModInverse(mod, num, 0, 1)
	return inverse % mod
}

func recurseModInverse(num1 int, num2 int, p0 int, p1 int) int {
	if num2 == 0 {
		return 0
	} else if num2 == 1 {
		return p1
	}

	p0, p1 = p1, (p0 - p1*(num1/num2))
	//  print(num1, num1//num2, num1 % num2, p1)
	num1, num2 = num2, num1%num2
	return recurseModInverse(num1, num2, p0, p1)
}
