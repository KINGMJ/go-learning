package remove_duplicates

func RemoveDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	// 双指针
	slow, fast := 1, 1
	for ; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
