package count_inversions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/daviskoh/algorithms_specialization/courses/01/assignments/02_count_inversions/count_inversions"
)

var _ = Describe("CountInversions", func() {
	Context("CountSplitInversions()", func() {
		var a, b, merged []int
		var count int

		It("handles simple case", func() {
			a = []int{1}
			b = []int{2}
			merged, count = CountSplitInversions(a, b)
			Expect(merged).To(Equal([]int{1, 2}))
			Expect(count).To(Equal(0))

			a = []int{2}
			b = []int{1}
			merged, count = CountSplitInversions(a, b)
			Expect(merged).To(Equal([]int{1, 2}))
			Expect(count).To(Equal(1))
		})

		It("handles longer slices", func() {
			a = []int{2, 3}
			b = []int{1, 4}
			merged, count = CountSplitInversions(a, b)
			Expect(merged).To(Equal([]int{1, 2, 3, 4}))
			Expect(count).To(Equal(2))

			a = []int{2, 4}
			b = []int{1, 3}
			merged, count = CountSplitInversions(a, b)
			Expect(merged).To(Equal([]int{1, 2, 3, 4}))
			Expect(count).To(Equal(3))
		})
	})

	Context("CountInversions()", func() {
		var count int
		var sorted []int

		It("handles base case", func() {
			for i := 0; i < 4; i++ {
				sorted, count = SortAndCount([]int{i})
				Expect(sorted).To(Equal([]int{i}))
				Expect(count).To(Equal(0))
			}
		})

		It("counts inversions", func() {
			// 2, 4, 1, 3, 5 => (2, 1), (4, 1), (4, 3) = 3
			sorted, count = SortAndCount([]int{2, 4, 1, 3, 5})
			Expect(sorted).To(Equal([]int{1, 2, 3, 4, 5}))
			Expect(count).To(Equal(3))

			// 3 4 6 1 2 5 => 7
			sorted, count = SortAndCount([]int{3, 4, 6, 1, 2, 5})
			Expect(sorted).To(Equal([]int{1, 2, 3, 4, 5, 6}))
			Expect(count).To(Equal(7))

			// 5 2 10 8 1 9 4 3 6 7
			sorted, count = SortAndCount([]int{5, 2, 10, 8, 1, 9, 4, 3, 6, 7})
			Expect(sorted).To(Equal([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
			Expect(count).To(Equal(22))

			// integers, _ := ioutil.ReadFile("./integers.txt")
			// fmt.Println(integers)
		})
	})
})
