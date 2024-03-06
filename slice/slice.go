package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := make([]int, 0)
	for {
		var x string
		fmt.Println("Enter an integer to sort or X to stop")
		fmt.Scan(&x)
		if x == "X" {
			return
		}
		input, err := strconv.Atoi(x)
		if err != nil {
			return
		}
		s = append(s, input)
		s = sort(s)
		fmt.Println(s)
	}
}

func sort(s []int) []int {
	sorted := s
	for i := 0; i < len(sorted); i++ {
		for j := 0; j < len(sorted)-1; j++ {
			if sorted[j] > sorted[j+1] {
				temp := sorted[j+1]
				sorted[j+1] = sorted[j]
				sorted[j] = temp
			}
		}
	}

	return sorted
}
