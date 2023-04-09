package util

func IsMember (arr []int, x int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == x {
			return true
		}
	}
	return false
}