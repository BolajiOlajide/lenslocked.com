package main

// import (
// 	"errors"
// 	"fmt"
// 	"html/template"
// 	"os"
// )

// // User structure for a user
// type User struct {
// 	Name string
// 	Age  int
// 	// anonymous struct
// 	// Meta struct {
// 	// 	Visits int
// 	// }
// 	Meta    UserMeta
// 	Bio     string
// 	RawHTML template.HTML
// }

// // UserMeta giving meta info about a user
// type UserMeta struct {
// 	Visits int
// }

// // ErrNotFound signifies when an item isn't found
// var ErrNotFound = errors.New("not found")

// func main() {
// 	t, err := template.ParseFiles("cmd/exp/hello.gohtml")

// 	if err != nil {
// 		panic(err)
// 	}

// 	user := User{
// 		Name: "Susan Smith",
// 		Age:  111,
// 		Meta: UserMeta{
// 			Visits: 4,
// 		},
// 		Bio:     `<script>alert("I am a user")</script>`,
// 		RawHTML: `<script>console.log("I AM A SCRIPT")</script>`,
// 	}

// 	err = t.Execute(os.Stdout, user)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = b()
// 	fmt.Println(errors.Is(err, ErrNotFound))
// 	unwrappedError := errors.Unwrap(err)
// 	fmt.Println(unwrappedError)
// 	// TODO determine if err is an ErrNotFound
// }

// func a() error {
// 	return ErrNotFound
// }

// func b() error {
// 	err := a()
// 	return fmt.Errorf("b: %v", err)
// }
