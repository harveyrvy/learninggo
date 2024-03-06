package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter a string:")
		s, _ = reader.ReadString('\n')
		fmt.Println(findIan(s))
	}
}

func findIan(s string) string {
	lowerS := strings.TrimSpace(strings.ToLower(s))
	if strings.HasPrefix(lowerS, "i") && strings.Contains(lowerS, "a") && strings.HasSuffix(lowerS, "n") {
		return "Found!"
	}
	return "Not Found!"
}
