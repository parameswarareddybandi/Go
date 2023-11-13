package main

import "fmt"

func main() {
	var rows, temp int
	fmt.Println("Enter rows :")
	fmt.Scan(&rows)
	for i := 0; i < rows; i++ {
		for k := 0; k <= i; k++ {
			temp = temp + 1
			fmt.Printf("%d ", temp)
		}
		fmt.Println()
	}
}
