package count_inversions_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCountInversions(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CountInversions Suite")
}
