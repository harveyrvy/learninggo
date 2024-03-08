package main

import (
	"fmt"
	"math"
)

func getInputs() (a, v_0, s_0 float64) {
	fmt.Println("enter acceleration")
	fmt.Scan(&a)
	fmt.Println("enter initial velocity")
	fmt.Scan(&v_0)
	fmt.Println("enter initial displacement")
	fmt.Scan(&s_0)

	return
}

func GenDisplaceFn(a, v_0, s_0 float64) func(t float64) float64 {
	fn := func(t float64) float64 {
		s := 0.5*a*math.Pow(t, 2) + v_0*t + s_0
		return s
	}
	return fn
}

func main() {
	var t float64
	a, v_0, s_0 := getInputs()
	fn := GenDisplaceFn(a, v_0, s_0)
	for {
		fmt.Println("enter time")
		fmt.Scan(&t)
		fmt.Printf("Displacement after time %f is %f \n", t, fn(t))
	}

}
