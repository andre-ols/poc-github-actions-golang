// name of the package
package user

// User is a struct of human data
type User struct {
	Age  int
	Name string
}

// Talk is a method of the User struct
func (receiver User) Talk() string {
	return "Every User Gets to Say Something"
}
