package piscine

func Abort(a, b, c, d, e int) int {
	nums := []int{a, b, c, d, e}
	sizen := len(nums)

	for i := 0; i < sizen; i++ {
		for j := 0; j < sizen-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums[sizen/2]
}
