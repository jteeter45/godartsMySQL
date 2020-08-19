package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func Newplayerset(w http.ResponseWriter, r *http.Request) {
	//ShowAllJson()
	posted = "N"
	tmpl.ExecuteTemplate(w, "Playerset", nil)
}

func Setnewplayers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Setplayers for All holdteam")
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM Player WHERE Team =?", holdteam)
	if err != nil {
		panic(err.Error())
	}
	ply := Player{}
	jply := Player_Small{}
	res := []Player{}
	jres := []Player_Small{}
	for selDB.Next() {
		var name, team, phone string
		err = selDB.Scan(&name, &team, &phone)
		if err != nil {
			panic(err.Error())
		}
		ply.Name = name
		res = append(res, ply)
		jply.Name_Json = name
		jres = append(jres, jply)

	}
	defer db.Close()
	fmt.Println("Endpoint Hit: bottom playersets")
	fmt.Println(res)
	fmt.Println(jres)

	html := `<!DOCTYPE html>
	<html>
	<body>
	<form action="http://localhost:8090/getselnewplayers" method="post" enctype="multipart/form-data">
	<p>Select a player</p>
	<select name="selplayers"> 
	{{range $key, $value := .}}
	<option value="{{ $value }}">{{ $value }}</option>
	{{end}}


	</select>
	<input type="submit" name="submit" value="Submit">

	</body>
	</html>`

	dropdownTemplate, err := template.New("dropdownexample").Parse(string(html))
	if err != nil {
		panic(err)
	}

	// populate dropdown with jres
	dropdownTemplate.Execute(w, jres)
}

func GetSelNewPlayers(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(200000) // grab the multipart form

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	//fmt.Fprintln(w, r.FormValue("selplayers"))
	holdplayer = r.FormValue("selplayers")
	//knock off the {}
	holdplayer = (strings.TrimRight(holdplayer, "}"))
	holdplayer = (strings.TrimPrefix(holdplayer, "{"))
	fmt.Println("GetSelPlayers,holdplayer: ", holdplayer)
	db := dbConn()
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	var player, team string
	var week int
	//err = db.QueryRow("SELECT Team, Week, Player  FROM Player_Stats where Week=? && Team =? && Player =?", holdweeknum, holdteam, holdplayer).Scan(&team, &week, &player )
	row := db.QueryRow("SELECT Week,Team,Player FROM Player_Stats where Week=? && Team =? && Player =?", holdweeknum, holdteam, holdplayer)

	switch err := row.Scan(&week, &team, &player); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		defer db.Close()
		http.Redirect(w, r, "/new", 301)
		//tmpl.ExecuteTemplate(w, "New", nil)
	case nil:
		defer db.Close()
		http.Redirect(w, r, "/", 301)
	default:
		defer db.Close()
		fmt.Fprintln(w, err)
	}
}
