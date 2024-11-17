package main

import (
	"fmt"
)

func CountPositivesSumNegatives(numbers []int) []int {
	var res []int
	var pos_count int
	var neg_sum int
	for _, val := range numbers {
		if val > 0 {
			pos_count += 1
		} else if val < 0 {
			neg_sum += val
		}
	}
	res = append(res, pos_count, neg_sum)
	return res
}

func main() {
	teste := []int{-1, 2, 3, -4, 5}

	fmt.Println(CountPositivesSumNegatives(teste))
}
