package sort

func ChoosePivotIndex(a []int) int {
	length := len(a)

	first := 0

	middle := length / 2
	if length%2 == 0 {
		middle -= 1
	}

	last := length - 1

	sortedByValues, _ := Sort([]int{a[first], a[middle], a[last]}, func(a []int) int {
		return 0
	})

	i := first
	switch sortedByValues[1] {
	case a[middle]:
		return middle
	case a[last]:
		return last
	}

	return i
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

func Sort(a []int, choosePivotIndex func(a []int) int) ([]int, int) {
	if len(a) < 2 {
		return a, 0
	}

	newPivotIndex := Partition(a, choosePivotIndex(a))

	comparisons := len(a) - 1

	left := a[:newPivotIndex]
	_, leftComparisons := Sort(left, choosePivotIndex)
	comparisons += leftComparisons

	right := a[newPivotIndex+1:]
	_, rightComparisons := Sort(right, choosePivotIndex)
	comparisons += rightComparisons

	return a, comparisons
}
