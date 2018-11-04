package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/daviskoh/algorithms_specialization/courses/01/assignments/02_count_inversions/count_inversions"
	// "github.com/daviskoh/algorithms_specialization/courses/01/assignments/02_count_inversions/count_inversions"
)

func readFile(fname string) (nums []int, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
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
	// fmt.Println(readFile("test.txt"))

	integers, _ := readFile("integers.txt")
	// fmt.Println(integers[0])

	_, count := count_inversions.SortAndCount(integers)
	fmt.Println(count)
}
