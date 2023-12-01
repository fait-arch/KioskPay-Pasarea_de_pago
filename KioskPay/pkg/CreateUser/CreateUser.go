package CreateUser

import (
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	Name  string
	Email string
	Age   int
}

func SubmitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	age, err := strconv.Atoi(r.FormValue("age"))
	if err != nil {
		http.Error(w, "Age must be a number", http.StatusBadRequest)
		return
	}

	user := User{
		Name:  name,
		Email: email,
		Age:   age,
	}

	fmt.Printf("%+v\n", user)
}
