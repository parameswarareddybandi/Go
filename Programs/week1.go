// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main

import "fmt"

func GCD(num1 int, num2 int) int {
	if num1 == 0 {
		return num2
	}
	if num2 == 0 {
		return num1
	}
	if num1 == num2 {
		return num1
	}
	if num1 > num2 {
		return GCD(num1-num2, num2)
	}
	return GCD(num1, num2-num1)
}

func main() {

	var num1, num2 int
	fmt.Println("Enter the numbers num1 and num2 : ")
	fmt.Scan(&num1, &num2)
	var gcd int = GCD(num1, num2)
	fmt.Println("GCD of", num1, "and", num2, "is:", gcd)
	var lcm int = (num1 * num2) / gcd
	fmt.Println("LCM of", num1, "and", num2, "is:", lcm)

}
