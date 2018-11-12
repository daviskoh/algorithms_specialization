package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/daviskoh/algorithms_specialization/courses/01/assignments/03_quicksort/sort"
)

func readFile(fname string) (nums []int, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\r\n")
	// Assign cap to avoid resize on every append.
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}

func main() {
	integers, _ := readFile("integers.txt")

	_, count := sort.Sort(integers, sort.ChoosePivotIndex)
	fmt.Println(count)
}
