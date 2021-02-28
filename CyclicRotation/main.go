package main

func Solution(A []int, K int) []int {

	if len(A) == 0 {
		return A
	}

	result := make([]int, len(A))

	for i, v := range A {
		newIdx := (i + K) % len(A)
		result[newIdx] = v
	}

	return result
}
