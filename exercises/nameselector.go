package main

import "math/rand"

var names = []string{"Alice", "Bob", "Charlie"}

func NameSelector() string {
	i := rand.Intn(len(names))
	return names[i]
}
