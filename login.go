package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	posted = "N"
	fmt.Println("I got to loginPage")
	tmpl.ExecuteTemplate(w, "Login", nil)

}

func Actuallogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		var databaseUsername string
		var databasePassword string
		var errfound int

		db := dbConn()
		err := db.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&databaseUsername, &databasePassword)
		if err != nil {
			fmt.Println("Server error...", err)
			//http.Error(w, "Server error, unable to find your account.", 500)
			errfound = 500
			//panic(err.Error())
		}

		err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
		if err != nil {
			fmt.Println("Server error..", err)
			//http.Error(w, "Server error, hashed password not found.", 500)
			errfound = 500
			//panic(err.Error())
		}

		if errfound == 500 {
			tmpl.ExecuteTemplate(w, "Login", nil)
		} else {
			defer db.Close()
			http.Redirect(w, r, "/", 301)
		}
	}
	//http.Redirect(w, r, "/", 301)
}
