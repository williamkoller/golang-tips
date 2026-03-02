package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	println(result[0], result[1])
}