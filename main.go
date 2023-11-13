package main

import (
	parent "family/father"
	child "family/father/son"
	"fmt"
)

func main() {
	f := new(parent.Father)
	fmt.Println(f.Data("B Jayasudhan Reddy"))

	c := new(child.Son)
	fmt.Println(c.Data("B Parameswara Reddy"))
}
