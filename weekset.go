package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Schedule struct {
	Week    int    `json:"week"`
	Date    string `json:"date"`
	Venue1  string `json:"venue1"`
	Team1   string `json:"team1"`
	Team2   string `json:"team2"`
	Venue2  string `json:"venue2"`
	Team3   string `json:"team3"`
	Team4   string `json:"team4"`
	Venue3  string `json:"venue3"`
	Team5   string `json:"team5"`
	Team6   string `json:"team6"`
	Venue4  string `json:"venue4"`
	Team7   string `json:"team7"`
	Teambye string `json:"teambye"`
}
type Schedule_Small struct {
	Week_Json int `json:"week"`
	//Date_Json string `json:"date"`
}

func Weekset(w http.ResponseWriter, r *http.Request) {
	//ShowAllJson()
	holdweek = ""
	holdweeknum = 0
	holdplayer = ""
	holdteam = ""
	posted = "N"
	fmt.Println("Endpoint Hit: Weekset")
	tmpl.ExecuteTemplate(w, "Weekset", nil)
}

func Schedules(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: Schedules")
	db := dbConn()

	//selDB, err := db.Query("SELECT distinct week, date FROM Schedule")
	selDB, err := db.Query("SELECT Week FROM Schedule")
	if err != nil {
		panic(err.Error())
	}
	ply := Schedule{}
	jply := Schedule_Small{}
	res := []Schedule{}
	jres := []Schedule_Small{}
	for selDB.Next() {
		//var date string
		var week int
		err = selDB.Scan(&week)
		if err != nil {
			panic(err.Error())
		}
		ply.Week = week
		res = append(res, ply)
		jply.Week_Json = week
		jres = append(jres, jply)

	}
	db.Close()
	fmt.Println("Endpoint Hit: bottom Schedules")
	fmt.Println(res)
	fmt.Println(jres)

	html := `<!DOCTYPE html>
<html>
<body>
<form action="http://localhost:8090/getselweek" method="post" enctype="multipart/form-data">
<p>Select a week</p>
<select name="selweek"> 
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

func GetSelWeek(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint GetSelWeek")
	err := r.ParseMultipartForm(200000) // grab the multipart form

	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	//fmt.Fprintln(w, r.FormValue("selweek"))
	holdweek = r.FormValue("selweek")
	//knock off the {}
	holdweek = (strings.TrimRight(holdweek, "}"))
	holdweek = (strings.TrimPrefix(holdweek, "{"))
	fmt.Println(holdweek)
	holdweeknum, err = strconv.Atoi(holdweek)
	fmt.Println("holdweeknum: ", holdweeknum)

	if holdweeknum == 0 {
		http.Redirect(w, r, "/", 301)
	} else {
		http.Redirect(w, r, "/teamset", 301)
	}
}
