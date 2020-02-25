func nextPermutation(nums []int)  {
	i, n := 1, len(nums)
	// index of swap nums
	out, in := -1, -1
	for ; i < n; i++ {
		if nums[i-1] < nums[i] {
			out, in = i-1, i
		} else if in!=-1 && nums[i]>nums[out] && nums[i]<=nums[in] {
			// find the smallest number greater than nums[out] and its index should be closes to end
			in = i
		}
	}
	if out == -1 {
		reverse(nums, 0, n-1)
	} else {
		nums[out], nums[in] = nums[in], nums[out]
		reverse(nums, out+1, n-1)
	}
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}