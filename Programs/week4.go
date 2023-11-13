package main

import (
	"fmt"
	"math"
)

func main() {
	var nums [10]float64
	var mean, sum, sd float64
	fmt.Println("******* Enter 10 Elements *******")
	for i := 0; i < 10; i++ {
		fmt.Scan(&nums[i])
		sum += nums[i]
	}
	mean = sum / 10
	for i := 0; i < 10; i++ {
		sd += math.Pow(nums[i]-mean, 2)
	}
	sd = math.Sqrt(sd / 10)
	fmt.Println("Standard Deviation of 10 elements is", sd)

}
