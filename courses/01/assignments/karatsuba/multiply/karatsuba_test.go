package multiply_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/daviskoh/algorithms_specialization/courses/01/assignments/karatsuba/multiply"
)

var _ = Describe("Karatsuba", func() {
	var x, y string

	// XContext("Pad()", func() {
	// 	It("pads left", func() {
	// 		Expect(Pad("1", 2, true)).To(Equal("001"))
	// 		Expect(Pad("1", 10, true)).To(Equal("00000000001"))
	// 	})
	//
	// 	It("pads right", func() {
	// 		Expect(Pad("1", 2, false)).To(Equal("100"))
	// 		Expect(Pad("1", 10, false)).To(Equal("10000000000"))
	// 	})
	// })

	Context("Karatsuba()", func() {
		It("handles simple cases", func() {
			x = "12"
			y = "1"
			Expect(Karatsuba(x, y)).To(Equal("12"))
		})

		It("pads digits correctly", func() {
			x = "123112"
			y = "1231"
			Expect(Karatsuba(x, y)).To(Equal("151550872"))

			x = "12244482"
			y = "1222"
			Expect(Karatsuba(x, y)).To(Equal("14962757004"))

			x = "7771829"
			y = "77718291231"
			Expect(Karatsuba(x, y)).To(Equal("604013269619531499"))

			x = "314159265358979323846264338327"
			y = "2718281828459045235360287471352662"
			Expect(Karatsuba(x, y)).To(Equal("853973422267356706546355086952074141385326159404594475100076474"))
		})

		It("does thangs...", func() {
			x = "2345"
			y = "123"
			Expect(Karatsuba(x, y)).To(Equal("288435"))

			x = "1"
			y = "123"
			Expect(Karatsuba(x, y)).To(Equal("123"))

			x = "12345"
			y = "123"
			Expect(Karatsuba(x, y)).To(Equal("1518435"))

			x = "110"
			y = "187"
			Expect(Karatsuba(x, y)).To(Equal("20570"))

			x = "4466"
			y = "13354"
			Expect(Karatsuba(x, y)).To(Equal("59638964"))

			x = "11223344"
			y = "55667788"
			Expect(Karatsuba(x, y)).To(Equal("624778734443072"))

			x = "12345678"
			y = "12345678"
			Expect(Karatsuba(x, y)).To(Equal("152415765279684"))

			x = "99999999"
			y = "99999999"
			Expect(Karatsuba(x, y)).To(Equal("9999999800000001"))

			x = "3141592653589793238462643383279502884197169399375105820974944592"
			y = "2718281828459045235360287471352662497757247093699959574966967627"
			Expect(Karatsuba(x, y)).To(Equal("8539734222673567065463550869546574495034888535765114961879601127067743044893204848617875072216249073013374895871952806582723184"))
		})
	})
})
