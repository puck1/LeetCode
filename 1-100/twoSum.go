func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, value := range nums {
		if j, ok := m[target - value]; ok {
			return []int{j, i}
		}
		m[value] = i
	}
	return []int{}
}