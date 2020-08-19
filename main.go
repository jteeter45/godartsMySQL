package main

import (
	"log"
	"net/http"

	//"text/template"
	"html/template"

	_ "github.com/go-sql-driver/mysql"
)

type Users struct {
	username string
	password string
	email    string
	uteam    string
}

type Player_Stats struct {
	Id                    int
	Week                  int
	Team                  string
	Player                string
	Five_singles_wldnp    string
	Five_singles_win      int
	Five_singles_loss     int
	Cricket_singles_wldnp string
	Cricket_singles_win   int
	Cricket_singles_loss  int
	Five_doubles_wldnp    string
	Five_doubles_win      int
	Five_doubles_loss     int
	Cricket_doubles_wldnp string
	Cricket_doubles_win   int
	Cricket_doubles_loss  int
	PlusOne               int
	TonEighty             int
	NineHtr               int
	MVP                   int
	HighOut               int
	SixBulls              int
}

var holdweek = ""
var holdweeknum = 0
var holdteam = ""
var holdplayer = ""
var posted = "N"
var secteam = ""

var tmpl = template.Must(template.ParseGlob("forms/*.tmpl"))

func handlerICon(w http.ResponseWriter, r *http.Request) {
}

func main() {
	log.Println("Server started on: http://localhost:8090")
	mux := http.NewServeMux()
	mux.HandleFunc("/favicon.ico", handlerICon)
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/show", Show)
	mux.HandleFunc("/new", New)
	mux.HandleFunc("/edit", Edit)
	mux.HandleFunc("/insert", Insert)
	mux.HandleFunc("/update", Update)
	mux.HandleFunc("/delete", Delete)
	mux.HandleFunc("/schedules", Schedules)
	mux.HandleFunc("/getselweek", GetSelWeek)
	mux.HandleFunc("/weekset", Weekset)
	mux.HandleFunc("/setteams", Setteams)
	mux.HandleFunc("/getselteam", GetSelTeam)
	mux.HandleFunc("/teamset", Teamset)
	mux.HandleFunc("/setplayers", Setplayers)
	mux.HandleFunc("/getselplayers", GetSelPlayers)
	mux.HandleFunc("/playerset", Playerset)
	mux.HandleFunc("/setnewplayers", Setnewplayers)
	mux.HandleFunc("/getselnewplayers", GetSelNewPlayers)
	mux.HandleFunc("/newplayerset", Newplayerset)
	mux.HandleFunc("/signup", Signup)
	mux.HandleFunc("/actualsignup", Actualsignup)
	mux.HandleFunc("/login", Login)
	mux.HandleFunc("/actuallogin", Actuallogin)
	http.ListenAndServe(":8090", mux)

	//http.ListenAndServe(":8080", nil)
}
