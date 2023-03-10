// name of the package
package main

// fmt is responsible for formatting
import (
	"fmt"

	"github.com/andre-ols/poc-github-actions-golang/user"
)

func main() {
	// human is an initialization of the User struct
	human := user.User{
		Age:  0,
		Name: "person",
	}

	fmt.Println(human.Talk())
}
