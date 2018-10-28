package multiply

import (
	"math/big"
)

// assume len(x) == len(y)

// (10*n)*ac + (10^(n/2))*(ad + bc) + bd

func Karatsuba(x, y *big.Int) *big.Int {
	a, _ := new(big.Int).SetString(x.String()[0:2], 10)
	b, _ := new(big.Int).SetString(x.String()[2:4], 10)

	c, _ := new(big.Int).SetString(y.String()[0:2], 10)
	d, _ := new(big.Int).SetString(y.String()[2:4], 10)

	step1 := new(big.Int).Mul(a, c)
	step2 := new(big.Int).Mul(b, d)
	step3 := new(big.Int).Mul(a.Add(a, b), c.Add(c, d))

	step3copy := step3
	step4sub1 := step3copy.Sub(step3copy, step2)
	step4 := step4sub1.Sub(step4sub1, step1)

	left, _ := new(big.Int).SetString(step1.String()+"0000", 10)
	middle, _ := new(big.Int).SetString(step2.String(), 10)
	right, _ := new(big.Int).SetString(step4.String()+"00", 10)

	return middle.Add(middle, left.Add(left, right))
}
