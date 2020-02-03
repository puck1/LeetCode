func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	if n <= 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	ans := [][]int{}
	for i := 0; i < n-3; {
		if nums[i] + nums[n-3] + nums[n-2] + nums[n-1] < target {		// efficiency improvment *
			for i++; i < n-3 && nums[i] == nums[i-1]; i++ {}
			continue
		}
		for j := i+1; j < n-2; {										
			if nums[i] + nums[j] + nums[j+1] + nums[j+2] > target {		// efficiency improvment *
				break
			}
			if nums[i] + nums[j] + nums[n-2] + nums[n-1] < target {		// efficiency improvment *
				for j++; j < n-2 && nums[j] == nums[j-1]; j++ {}
				continue
			}
			left, right := j+1, n-1
			for left < right {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum < target {
					for left++; left < right && nums[left] == nums[left-1]; left++ {}
				} else if sum > target {
					for right--; left < right && nums[right] == nums[right+1]; right-- {}
				} else {
					ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {}
					for right--; left < right && nums[right] == nums[right+1]; right-- {}
				}
			}
			for j++; j < n-2 && nums[j] == nums[j-1]; j++ {}
		}
		for i++; i < n-3 && nums[i] == nums[i-1]; i++ {}
	}
	return ans
}