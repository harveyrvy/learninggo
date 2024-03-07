package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a name:")
	name, _ := reader.ReadString('\n')
	fmt.Println("Enter an address:")
	address, _ := reader.ReadString('\n')
	m := make(map[string]string)
	m["name"] = name
	m["address"] = address
	fmt.Println(m)
	// Marshall
	barr, _ := json.Marshal(m)
	// Create new map
	var m2 map[string]string
	// Unmarshall into new map and print
	json.Unmarshal(barr, &m2)
	fmt.Println(m2)
	return
}
