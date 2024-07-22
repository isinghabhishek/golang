package main

import "fmt"

// to run type in command prompt -- go run 6_if_else/ifelse.go

func main() {
	// age := 16

	// if age >= 18 {
	// 	fmt.Println("person is an adult")
	// } else {
	// 	fmt.Println("person is not an adult")
	// }

	// declaring variable
	// var age int

	// fmt.Println("Enter your age: ")
	// fmt.Scanln(&age)
	// if age >= 18 {
	// 	fmt.Println("Person is adult")
	// } else if age <= 12 {
	// 	fmt.Println("Person is not an adult")

	// } else {
	// 	fmt.Println("Person is kid")
	// }

	// var role = "admin"
	// var hasPermissions = false
	// if role == "admin" && hasPermissions {
	// 	fmt.Println("Yes")
	// }else if role == "" && hasPermissions {
	// 	fmt.Println("")
	// }

	// Features
	if age := 15; age >= 18 {
		fmt.Println("person is an adult", age)
	} else if age >= 12 {
		fmt.Println("person is teenager", age)
	}
}
