// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main

import "fmt"

func isPalindrome(str1 string) int {
	for i := 0; i < len(str1)/2; i++ {
		j := len(str1) - i - 1
		if str1[i] != str1[j] {
			return 0
		}
	}
	return 1
}

func main() {
	var str1 string
	fmt.Println("Enter a string : ")
	fmt.Scan(&str1)
	if isPalindrome(str1) == 1 {
		fmt.Println(str1, "is palindrome.")
	} else {
		fmt.Println(str1, "is not a palindrome.")
	}
}
