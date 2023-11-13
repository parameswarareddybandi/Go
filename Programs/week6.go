// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main
import "fmt"

func main() {
    var str1 string
    fmt.Println("Enter the first string : ")
    fmt.Scan(&str1)
    var str2 string
    fmt.Println("Enter the second string : ")
    fmt.Scan(&str2)
    var str string = str1 + str2
    fmt.Println(str)
}