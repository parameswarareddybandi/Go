package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ContainsAny("Paramesh", "mn"))
	fmt.Println(strings.ContainsAny("Paramesh", "p"))
	fmt.Println(strings.Contains("CMRIT", "CMR"))
	fmt.Println(strings.Contains("CMRIT", "cmr"))
	fmt.Println(strings.Count("malayalam", "a"))
	fmt.Println(strings.Count("mathematics", "m"))
	fmt.Println(strings.EqualFold("Cat", "cAt"))
	fmt.Println(strings.EqualFold("For", "Form"))
}
