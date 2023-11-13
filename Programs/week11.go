package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{50, 90, 30, 60, 40}
	fmt.Println("Interger Reverse Sort : ")
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	fmt.Println(nums)
	fmt.Println()

	text := []string{"Japan", "UK", "Uk", "INDIA", "Canada", "Germany"}
	fmt.Println("Strings Reverse Sort : ")
	sort.Sort(sort.Reverse(sort.StringSlice(text)))
	fmt.Println(text)
}
