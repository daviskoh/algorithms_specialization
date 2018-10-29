package multiply

import (
	"fmt"
	"math/big"
	"strconv"
)

func Pad(n string, by int, isLeft bool) string {
	nCopy := n

	for i := 0; i < by; i++ {
		if isLeft {
			nCopy = fmt.Sprintf("0%s", nCopy)
		} else {
			nCopy = fmt.Sprintf("%s0", nCopy)
		}

	}

	return nCopy
}

func Karatsuba(x, y string) string {
	xCopy := x
	yCopy := y

	xLen := len(x)
	yLen := len(y)

	if xLen == 1 && yLen == 1 {
		xInt, _ := strconv.Atoi(x)
		yInt, _ := strconv.Atoi(y)
		resp := strconv.Itoa(xInt * yInt)
		return resp
	}

	if xLen < yLen {
		xCopy = Pad(xCopy, yLen-xLen, true)
	} else if yLen < xLen {
		yCopy = Pad(yCopy, xLen-yLen, true)
	}

	digits := len(xCopy)
	middle := digits / 2
	if digits%2 != 0 {
		middle++
	}

	a := xCopy[:middle]
	b := xCopy[middle:]

	c := yCopy[:middle]
	d := yCopy[middle:]

	ac, _ := new(big.Int).SetString(Karatsuba(a, c), 10)
	bd, _ := new(big.Int).SetString(Karatsuba(b, d), 10)

	aInt, _ := new(big.Int).SetString(a, 10)
	bInt, _ := new(big.Int).SetString(b, 10)
	cInt, _ := new(big.Int).SetString(c, 10)
	dInt, _ := new(big.Int).SetString(d, 10)

	aPlusB := new(big.Int).Add(aInt, bInt)
	cPlusD := new(big.Int).Add(cInt, dInt)
	abTimescd := new(big.Int)
	abTimescd.SetString(Karatsuba(aPlusB.String(), cPlusD.String()), 10)

	gaussSub := new(big.Int).Sub(abTimescd, ac)
	gauss := new(big.Int).Sub(gaussSub, bd)

	left, _ := new(big.Int).SetString(Pad(ac.String(), (digits-middle)*2, false), 10)
	mid, _ := new(big.Int).SetString(Pad(gauss.String(), digits-middle, false), 10)

	right := bd

	respSub := new(big.Int).Add(left, mid)

	resp := new(big.Int).Add(respSub, right).String()
	return resp
}
