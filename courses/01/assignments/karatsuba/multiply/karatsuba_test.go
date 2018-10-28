package multiply_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"math/big"

	. "github.com/daviskoh/algorithms_specialization/courses/01/assignments/karatsuba/multiply"
)

var _ = Describe("Karatsuba", func() {
	It("handles base case", func() {
		var x, y big.Int

		x.SetString("5678", 10)
		y.SetString("1234", 10)
		Expect(Karatsuba(&x, &y).String()).To(Equal("7006652"))

		x.SetString("2233", 10)
		y.SetString("4455", 10)
		Expect(Karatsuba(&x, &y).String()).To(Equal("9948015"))

		x.SetString("1423", 10)
		y.SetString("5867", 10)
		Expect(Karatsuba(&x, &y).String()).To(Equal("8348741"))
	})
})
