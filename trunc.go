package main

import "fmt"

func main() {
	for {
		var fp float64
		fmt.Printf("Enter a floating point number \n")
		fmt.Scan(&fp)
		fmt.Printf("%d \n", truncate(fp))
	}
}

func truncate(x float64) int {
	y := int(x)
	return y
}
