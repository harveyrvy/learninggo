package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				Swap(nums, j)

			}
		}
	}
}

// Swap item at index with one AFTER
func Swap(nums []int, index int) {
	temp := nums[index+1]
	nums[index+1] = nums[index]
	nums[index] = temp
}

func SplitInput(input string) []string {
	splitList := strings.Split(input, " ")
	return splitList

}

func convertToInt(inputs []string) []int {
	var intList []int
	for _, v := range inputs {
		if v != "" {
			intValue, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			intList = append(intList, intValue)
		}

	}
	return intList
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	inputs, err := reader.ReadString('\n')
	inputs = strings.TrimSuffix(inputs, "\r\n")
	if err != nil {
		panic(err)
	}
	return inputs
}

func main() {
	fmt.Println("Type in up to 10 integers for sorting:")

	inputs := readInput()

	splitList := SplitInput(inputs)

	intList := convertToInt(splitList)

	BubbleSort(intList)

	fmt.Println(intList)
}
