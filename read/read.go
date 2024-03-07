package main

import (
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

var filename string
var names []Name

func main() {
	dat := make([]byte, 50)
	fmt.Println("Enter the name of the file: ")
	fmt.Scan(&filename)
	dat, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
	}

	lines := strings.Split(string(dat), "\n")

	for _, line := range lines {
		n := new(Name)
		inputNames := strings.Split(line, " ")

		n.fname = inputNames[0]
		n.lname = inputNames[1]
		names = append(names, *n)
	}

	for _, n := range names {
		fmt.Println(n.fname, n.lname)
	}

}
