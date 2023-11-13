package main

import "fmt"

func main() {
	var rows int
	fmt.Println("Enter the no. of rows : ")
	fmt.Scan(&rows)
	for i := 1; i <= rows; i++ {
		for spaces := 1; spaces <= rows-i; spaces++ {
			fmt.Print("  ")
		}
		size := 2*i - 1
		first := (size / 2) + 1
		second := size - first
		k := i
		for j := 1; j <= first; j++ {
			fmt.Printf(" %d", k)
			k++
		}
		k = i + (size / 2) - 1
		for j := 1; j <= second; j++ {
			fmt.Printf(" %d", k)
			k--
		}
		fmt.Println()
	}
}
