package remove_duplicates

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RemoveDuplicates", func() {
	It("test1", func() {
		var nums = []int{1, 1, 2}
		var length = RemoveDuplicates(nums)
		Expect(length).To(Equal(2))
		Expect(nums[:length]).To(Equal([]int{1, 2}))
	})

	It("test2", func() {
		var nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		var length = RemoveDuplicates(nums)
		Expect(length).To(Equal(5))
		Expect(nums[:length]).To(Equal([]int{0, 1, 2, 3, 4}))
	})
})
