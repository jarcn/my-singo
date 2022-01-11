package test

import "testing"

func TestSf1(t *testing.T) {
	arr := []int{3, 4, 56, 6, 7, 9}
	sum := 10
	t.Log(sf1(arr, sum))
}

//在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标。
func sf1(arr []int, sum int) []int {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		another := sum - arr[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[arr[i]] = i
	}
	return nil

}
