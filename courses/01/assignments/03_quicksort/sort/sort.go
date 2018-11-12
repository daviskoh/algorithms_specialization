package sort

import (
	"sort"
)

func ChoosePivotIndex(a []int) int {
	length := len(a)

	first := 0

	middle := length / 2
	if length%2 == 0 {
		middle -= 1
	}

	last := length - 1

	// going to cheat a little here
	sortedByValues := []int{first, middle, last}
	sort.Slice(sortedByValues, func(i int, j int) bool {
		return a[sortedByValues[i]] < a[sortedByValues[j]]
	})

	return sortedByValues[1]
}

func Partition(a []int, pivotIndex int) int {
	length := len(a)

	if length == 1 {
		return 0
	}

	a[pivotIndex], a[0] = a[0], a[pivotIndex]

	pivot := a[0]
	i := 1
	for j := i; j < length; j++ {
		if a[j] < pivot {
			a[j], a[i] = a[i], a[j]
			i++
		}
	}

	a[0], a[i-1] = a[i-1], a[0]

	return i - 1
}

func Sort(a []int) ([]int, int) {
	if len(a) < 2 {
		return a, 0
	}

	newPivotIndex := Partition(a, ChoosePivotIndex(a))

	comparisons := len(a) - 1

	left := a[:newPivotIndex]
	_, leftComparisons := Sort(left)
	comparisons += leftComparisons

	right := a[newPivotIndex+1:]
	_, rightComparisons := Sort(right)
	comparisons += rightComparisons

	return a, comparisons
}
