package count_inversions

// takes 2 sorted slices
func CountSplitInversions(a, b []int) ([]int, int) {
	aLength := len(a)
	bLength := len(b)

	n := len(a) + len(b)

	merged := make([]int, 0, n)
	inversions := 0

	i, j := 0, 0

	for k := 0; k < n; k++ {
		if i == aLength {
			merged = append(merged, b[j:]...)
			return merged, inversions
		}

		if j == bLength {
			merged = append(merged, a[i:]...)
			return merged, inversions
		}

		if a[i] < b[j] {
			merged = append(merged, a[i])
			i++
		} else {
			merged = append(merged, b[j])
			j++

			// inc total by elements remaining in a
			inversions += (aLength - i)
		}
	}

	return merged, inversions
}

func SortAndCount(numbers []int) ([]int, int) {
	n := len(numbers)

	if n == 1 {
		return numbers, 0
	}

	a, x := SortAndCount(numbers[:n/2])
	b, y := SortAndCount(numbers[n/2:])
	merged, z := CountSplitInversions(a, b)

	return merged, (x + y + z)
}
