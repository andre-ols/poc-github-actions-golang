// name of the package
package main

// fmt is responsible for formatting
import (
	"fmt"
)

// User is a struct of human data
type User struct {
	Age  int
	Name string
}

func main() {
	// human is an initialization of the User struct
	human := User{
		Age:  0,
		Name: "person",
	}

	fmt.Println(human.Talk())
}

// Talk is a method of the User struct
func (receiver User) Talk() string {
	return "Every User Gets to Say Something"
}
