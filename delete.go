package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	fmt.Println("In Delete ID is :", nId)
	delForm, err := db.Prepare("DELETE FROM Player_Stats WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(nId)
	log.Println("DELETE", nId)
	defer db.Close()
	posted = "N"
	http.Redirect(w, r, "/", 301)
}
