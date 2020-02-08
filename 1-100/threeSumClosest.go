// import (
// 	"sort"
// 	"math"
// )
func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	var ans int
	if n <= 3 {
		for _, num := range nums {
			ans += num
		}
		return ans
	}

	sort.Ints(nums)
	ans = nums[0] + nums[1] + nums[2]
	for i := 0; i < n-2; {
		lastDIff := math.MaxFloat64
		for left, right := i+1, n-1; left < right; {
			sum := nums[i] + nums[left] + nums[right]
			if diff := math.Abs(float64(target - sum)); diff >= lastDIff {
				break
			}
			if sum == target {
				return target
			}
			if math.Abs(float64(sum - target)) < math.Abs(float64(ans - target)) {
				ans = sum
			}
			if sum < target {
				for left++; left < right && nums[left] == nums[left-1]; left++ {}
			} else {
				for right--; left < right && nums[right] == nums[right+1]; right-- {}
			}
		}
		for i++; i < n-2 && nums[i] == nums[i-1]; i++ {}
	}
	return ans
}