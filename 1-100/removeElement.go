// 1.
func removeElement(nums []int, val int) int {
	cnt := 0
	for i := range nums {
		if nums[i] == val {
			cnt++
		} else {
			nums[i-cnt] = nums[i]
		}
	}
	return len(nums)-cnt
}

// 2.
func removeElement(nums []int, val int) int {
	cnt := 0
	for i := range nums {
		if nums[i] != val {
			nums[cnt] = nums[i]
			cnt++
		}
	}
	return cnt
}

// 3. since elements left in nums array can be arbitrary, 
// when elements to remove are rare, method 3 is more efficient *
func removeElement(nums []int, val int) int {
	i := 0
	n := len(nums)
	for i < n {
		if nums[i] == val {
			nums[i], nums[n-1] = nums[n-1], nums[i]		//swap
			n--
		} else {
			i++
		}
	}
	return n
}
