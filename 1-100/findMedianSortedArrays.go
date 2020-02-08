// *
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	if l1 > l2 {
		nums1, nums2 = nums2, nums1
		l1, l2 = l2, l1
	}
	imin ,imax := 0, l1
	var left_max, right_min int
	for imin <= imax {
		i := (imin + imax) / 2
		j := (l1 + l2 + 1) / 2 - i
		if i < imax && nums2[j - 1] > nums1[i] {
			imin = i + 1
		} else if i > imin && nums1[i - 1] > nums2[j] {
			imax = i - 1
		} else {
			if i == 0 {
				left_max = nums2[j - 1]
			} else if j == 0 {
				left_max = nums1[i - 1]
			} else if nums1[i - 1] > nums2[j - 1] {
				left_max = nums1[i - 1]
			} else {
				left_max = nums2[j - 1]
			}
			if (l1 + l2) % 2 == 1 {
				return float64(left_max)
			}
			if i == l1 {
				right_min = nums2[j]
			} else if j == l2 {
				right_min = nums1[i]
			} else if nums1[i] < nums2[j] {
				right_min = nums1[i]
			} else {
				right_min = nums2[j]
			}
			return float64(left_max + right_min) / 2
		}
	}
	return 0.0
}