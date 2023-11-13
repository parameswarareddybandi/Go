package main

import "fmt"

func removeDuplicates(s []string) []string {
	bucket := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := bucket[str]; !ok {
			bucket[str] = true
			result = append(result, str)
		}
	}
	return result
}

func main() {
	array := []string{"cseb", "section", "2024", "batch", "cseb", "section", "cmrit"}
	fmt.Println("Original Array : ", array)

	result := removeDuplicates(array)
	fmt.Println("After Removing Duplicates :", result)

}
