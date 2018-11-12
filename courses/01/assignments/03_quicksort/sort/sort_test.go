package sort_test

import (
	"sort"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/daviskoh/algorithms_specialization/courses/01/assignments/03_quicksort/sort"
)

var _ = Describe("QuickSort", func() {
	Context("ChoosePivotIndex()", func() {
		var a []int

		It("handles simple case odd length case", func() {
			a = []int{1, 2, 3}
			Expect(ChoosePivotIndex(a)).To(Equal(1))

			a = []int{1, 2, 3, 4, 6}
			Expect(ChoosePivotIndex(a)).To(Equal(2))

			a = []int{6, 4, 7, 4, 2}
			Expect(ChoosePivotIndex(a)).To(Equal(0))
		})

		It("handles even lengths", func() {
			a = []int{1, 2, 3, 4}
			Expect(ChoosePivotIndex(a)).To(Equal(1))

			a = []int{1, 2, 3, 4, 6, 7}
			Expect(ChoosePivotIndex(a)).To(Equal(2))

			a = []int{8, 2, 4, 5, 7, 1}
			Expect(ChoosePivotIndex(a)).To(Equal(2))
		})
	})

	Context("Partition()", func() {
		var a []int

		It("handles simple case", func() {
			a = []int{1, 2, 3}
			Partition(a, 0)
			Expect(a).To(Equal([]int{1, 2, 3}))
		})

		It("partitions around a selected pivot index", func() {
			a = []int{3, 1, 2}
			Partition(a, 2)
			Expect(a).To(Equal([]int{1, 2, 3}))
		})

		It("handles complex cases", func() {
			a = []int{1, 8, 2, 5, 3, 4, 7, 6}
			Partition(a, 4)
			Expect(a).To(Equal([]int{1, 2, 3, 5, 8, 4, 7, 6}))
		})
	})

	Context("Sort()", func() {
		var a, resp, sorted []int
		var comparisons int

		It("handles base case", func() {
			a = []int{1}
			resp, _ = Sort(a)
			Expect(resp).To(Equal([]int{1}))
		})

		It("handles simple case", func() {
			a = []int{2, 1}
			resp, _ = Sort(a)
			Expect(resp).To(Equal([]int{1, 2}))
		})

		It("handles complex cases", func() {
			a = []int{1, 3, 2}
			resp, _ = Sort(a)
			Expect(resp).To(Equal([]int{1, 2, 3}))

			a = []int{1, 7, 2, 5, 3, 4, 8, 6}
			resp, _ = Sort(a)
			sorted = make([]int, len(a))
			copy(sorted, a)
			sort.Slice(sorted, func(i int, j int) bool {
				return sorted[i] < sorted[j]
			})
			Expect(resp).To(Equal(sorted))

			a = []int{10, 5, 6, 4, 2, 7, 8, 9, 11, 14, 1}
			resp, _ = Sort(a)
			sorted = make([]int, len(a))
			copy(sorted, a)
			sort.Slice(sorted, func(i int, j int) bool {
				return sorted[i] < sorted[j]
			})
			Expect(resp).To(Equal(sorted))
		})

		XIt("returns comparisons count", func() {
			a = []int{1, 3, 2}
			_, comparisons = Sort(a)
			Expect(comparisons).To(Equal(3))

			a = []int{4, 5, 6, 3}
			_, comparisons = Sort(a)
			Expect(comparisons).To(Equal(4))
		})
	})
})
