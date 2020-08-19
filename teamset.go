package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Team struct {
	Team         string `json:"team"`
	Abbreviation string `json:"abbreviation"`
}

type Team_Small struct {
	Team_Json string `json:"team"`
}

func Teamset(w http.ResponseWriter, r *http.Request) {
	//ShowAllJson()
	posted = "N"
	tmpl.ExecuteTemplate(w, "Teamset", nil)
}

func Setteams(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Setteams")
	db := dbConn()

	selDB, err := db.Query("SELECT Team FROM Team")
	if err != nil {
		panic(err.Error())
	}
	ply := Team{}
	jply := Team_Small{}
	res := []Team{}
	jres := []Team_Small{}
	for selDB.Next() {
		var team string
		err = selDB.Scan(&team)
		if err != nil {
			panic(err.Error())
		}
		ply.Team = team
		res = append(res, ply)
		jply.Team_Json = team
		jres = append(jres, jply)

	}
	defer db.Close()
	fmt.Println("Endpoint Hit: bottom Setteams")
	fmt.Println(res)
	fmt.Println(jres)

	html := `<!DOCTYPE html>
<html>
<body>
<form action="http://localhost:8090/getselteam" method="post" enctype="multipart/form-data">
<p>Select a team</p>
<select name="selteam"> 
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
	//http.Redirect(w, r, "/", 301)
	//w.Write([]byte(fmt.Sprintf(html)))
}

func GetSelTeam(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint GetSelTeam")
	err := r.ParseMultipartForm(200000) // grab the multipart form

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	//fmt.Fprintln(w, r.FormValue("selteam"))
	holdteam = r.FormValue("selteam")
	//knock off the {}
	holdteam = (strings.TrimRight(holdteam, "}"))
	holdteam = (strings.TrimPrefix(holdteam, "{"))
	fmt.Println("holdteam: ", holdteam)

	if holdteam == "None" {
		http.Redirect(w, r, "/", 301)
	} else {
		http.Redirect(w, r, "playerset", 301)
	}

	//tmpl.ExecuteTemplate(w, "Playerset", nil)
}
