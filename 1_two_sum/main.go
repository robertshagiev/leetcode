package main

func twoSum(nums []int, target int) []int {
	//подход двойного цикла (brute force) - метод заключается в проверке каждой пары чисел в массиве
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
