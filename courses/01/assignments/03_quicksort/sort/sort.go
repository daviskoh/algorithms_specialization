package sort

func ChoosePivotIndex(a []int, l, r int) int {
	first := 0

	middle := r / 2
	if r%2 == 0 {
		middle -= 1
	}

	last := r - 1

	values := []int{a[first], a[middle], a[last]}
	sortedByValues, _ := Sort(values, 0, 3, func(a []int, l, r int) int {
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

func Partition(a []int, l, r int, pivotIndex int) int {
	length := len(a)

	if length == 1 {
		return 0
	}

	if pivotIndex != 0 {
		a[pivotIndex], a[0] = a[0], a[pivotIndex]
	}

	pivot := a[l]
	i := l + 1
	for j := l + 1; j < r; j++ {
		if a[j] < pivot {
			a[j], a[i] = a[i], a[j]
			i++
		}
	}

	a[l], a[i-1] = a[i-1], a[l]

	return i - 1
}

var i = 0

func Sort(a []int, l, r int, choosePivotIndex func(a []int, l, r int) int) ([]int, int) {
	if r-l < 2 {
		return a, 0
	}

	pivotIndex := choosePivotIndex(a, l, r)
	currentPivotPos := Partition(a, l, r, pivotIndex)

	comparisons := len(a) - 1

	_, leftComparisons := Sort(a, 0, currentPivotPos, choosePivotIndex)
	comparisons += leftComparisons

	_, rightComparisons := Sort(a, currentPivotPos+1, r, choosePivotIndex)
	comparisons += rightComparisons

	return a, comparisons
}
