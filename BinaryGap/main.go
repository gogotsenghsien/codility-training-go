package main

import (
	"math"
	"strconv"
)

func Solution(N int) int {
	// write your code in Go 1.4
	if N == 0 {
		return 0
	}

	binaryInt := strconv.FormatInt(int64(N), 2)
	result := 0
	i := 0

	for i < len(binaryInt) {
		cnt := 0

		if i > 0 && string(binaryInt[i]) == "0" && string(binaryInt[i-1]) == "1" {
			for i < len(binaryInt) && string(binaryInt[i]) != "1" {
				i++
				cnt++
			}

			if i < len(binaryInt) {
				result = int(math.Max(float64(result), float64(cnt)))
			}
		} else {
			i++
		}
	}

	return result
}
