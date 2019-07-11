package removeDuplicates

// 给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[j] {
			continue
		}
		j += 1
		if i == j {
			continue
		}
		nums[j] = nums[i]
	}
	return j + 1
}
