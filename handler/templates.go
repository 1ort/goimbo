package handler

//used in html templates for page list
func IntRange(n int) []int {
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = i
	}
	return res
}
