// 1.
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dupCnt := 0
	for i := 1; i < n; i++ {
		if nums[i] == nums[i-1] {
			dupCnt++
		} else {
			nums[i-dupCnt] = nums[i]
		}
	}
	return n-dupCnt
}

// 2.
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	cnt := 1
	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			nums[cnt] = nums[i]
			cnt++
		}
	}
	return cnt
}