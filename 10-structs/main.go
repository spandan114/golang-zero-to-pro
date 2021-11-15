package main

import "fmt"

func main() {
	//structs is similar like class
	//no inharitance,super,parent ,encapsulation etc...

	newUser := User{"spandan", "spandan@go.dev", true}
	fmt.Printf("%+v\n", newUser)
	fmt.Println(newUser.name)

	//slice of struct
	arrUser := []User{}
	arrUser = append(arrUser, User{"spandan", "spandan@go.dev", true})
	fmt.Println(arrUser)
}

type User struct {
	name       string
	email      string
	isVarified bool
}
