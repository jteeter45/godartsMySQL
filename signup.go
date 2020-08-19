package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I got to signupPage")
	posted = "N"
	tmpl.ExecuteTemplate(w, "Signup", nil)
}

func Actualsignup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db := dbConn()
		username := r.FormValue("username")
		password := r.FormValue("password")
		password1 := r.FormValue("password1")
		if password != password1 {
			http.Error(w, "Passwords not equal", 500)
			defer db.Close()
			http.Redirect(w, r, "/Login", 301)
		}
		email := r.FormValue("email")
		uteam := r.FormValue("uteam") //MAY NEED TO DO DROP DOWN TO ASSURE VALID ENTRY

		var user string

		err := db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)

		//err := db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)

		switch {
		case err == sql.ErrNoRows:
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
			if err != nil {
				fmt.Println("Server error, unable to create your account.", err)
				defer db.Close()
				http.Redirect(w, r, "/Login", 301)
			}

			fmt.Println("username: ", username)
			fmt.Println("hashedPassword: ", hashedPassword)
			fmt.Println("email: ", email)
			fmt.Println("uteam: ", uteam)

			_, err = db.Exec("INSERT INTO users(username, password, email, uteam) VALUES(?, ?, ?, ?)", username, hashedPassword, email, uteam)
			if err != nil {
				fmt.Println("Server error, unable to create your account.", err)
				defer db.Close()
				//http.Error(w, "Server error, unable to create your account.", 500)
				http.Redirect(w, r, "/Login", 301)
			}

			fmt.Println("User created!")
			http.Redirect(w, r, "/", 301)
		case err != nil:
			fmt.Println("Server error, unable to create your account.", err)
			defer db.Close()
			http.Redirect(w, r, "/Login", 301)
		default:
			defer db.Close()
			http.Redirect(w, r, "/", 301)
		}
	}

	//http.Redirect(w, r, "/", 301)
}
