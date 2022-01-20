package main

import (
	"html/template"
	"os"
)

// User structure for a user
type User struct {
	Name string
	Age  int
	// anonymous struct
	// Meta struct {
	// 	Visits int
	// }
	Meta    UserMeta
	Bio     string
	RawHTML template.HTML
}

// UserMeta giving meta info about a user
type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("cmd/exp/hello.gohtml")

	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Susan Smith",
		Age:  111,
		Meta: UserMeta{
			Visits: 4,
		},
		Bio:     `<script>alert("I am a user")</script>`,
		RawHTML: `<script>console.log("I AM A SCRIPT")</script>`,
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
