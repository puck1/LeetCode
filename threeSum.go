func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	for i := 0; i < len(nums) - 2; {
		toFound := map[int]int{}
		found := map[int]bool{}
		for j := i + 1; j < len(nums); j++ {
			if found[nums[j]] {		// skip dupulicate elem
				continue
			}
			if index, ok := toFound[-nums[i] - nums[j]]; ok {
				ans = append(ans, []int{nums[i], nums[index], nums[j]})
				found[nums[j]] = true		
			} else {
				toFound[nums[j]] = j
			}
		}
		for i++; i < len(nums) - 2 && nums[i] == nums[i-1]; i++ {}
	}
	return ans
}

// *
// import "sort"
func threeSum(nums []int) [][]int {
	n := len(nums)
	if n < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	ans := [][]int{}
	for i := 0; i < n - 2; {
		for left, right := i+1, n - 1; left < right; {
            sum := nums[i] + nums[left] + nums[right]
			if sum > 0 {
				for right--; left < right && nums[right] == nums[right+1]; right-- {}
			} else if sum < 0 {
				for left++; left < right && nums[left] == nums[left-1]; left++ {}
            } else {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for right--; left < right && nums[right] == nums[right+1]; right-- {}
				for left++; left < right && nums[left] == nums[left-1]; left++ {}                
            }
		}
		for i++; i < n - 2 && nums[i] == nums[i-1]; i++ {}
	}
	return ans
}