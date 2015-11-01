package main

import "math/rand"
import (
	"fmt"
	"time"
	"math"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var threshold = [...]float64{-math.MaxFloat64, -4, -3.5, -3, -2.5, -2, -1.5, -1, -0.5, 0, 0.5, 1, 1.5, 2, 2.5, 3, 3.5, 4, math.MaxFloat64}
	var count [18]int32
	var list [10000]float64
	for i := 0; i < 10000; i++ {
		list[i] = rand.NormFloat64()
	}

	for i := 0; i < 10000; i++ {
		for j := 0; j < len(count); j++ {
			if (threshold[j] < list[i] && list[i] < threshold[j + 1]) {
				count[j] = count[j] + 1
			}
		}
	}
	for j := 0; j < len(count); j++ {
		fmt.Print(threshold[j])
		fmt.Print("  ")
		fmt.Println(threshold[j + 1])
		fmt.Println(count[j])
	}
}
