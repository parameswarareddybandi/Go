package main

import "fmt"

func main() {
	var nums [20]int
	var avg, sum, size int
	fmt.Println("Enter size of array : ")
	fmt.Scan(&size)
	fmt.Println("Enter the elements of array : ")
	for i := 0; i < size; i++ {
		fmt.Scan(&nums[i])
		sum += nums[i]
	}
	avg = sum / size
	fmt.Println("Avergae :", avg)
}
