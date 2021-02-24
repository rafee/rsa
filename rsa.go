package rsa

func generateKeys() (int, int, int) {
	p, q := 34667, 44207
	N := p * q
	phiN := (p - 1) * (q - 1)

	// Very bad practice
	e := phiN - 1
	pk := modInverse(e, phiN)
	return N, e, pk
}

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
