package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sqweek/dialog"
)

func New(w http.ResponseWriter, r *http.Request) {
	//ShowAllJson()
	posted = "N"
	ok := dialog.Message("%s", "Do you remember your Week, Team and Player selections?").Title("Are you sure?").YesNo()
	if !ok {
		fmt.Println("they chose no")
		http.Redirect(w, r, "/", 301)
	}
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		//Player  := r.FormValue("player ")
		//Week := r.FormValue("week")
		//Team := r.FormValue("team")
		Week := holdweeknum
		Team := holdteam
		Player := holdplayer
		Five_singles_wldnp := r.FormValue("five_singles_wldnp")
		Cricket_singles_wldnp := r.FormValue("cricket_singles_wldnp")
		fmt.Println("New from form")
		Five_doubles_wldnp := r.FormValue("five_doubles_wldnp")
		Cricket_doubles_wldnp := r.FormValue("cricket_doubles_wldnp")
		PlusOne := r.FormValue("plusone")
		TonEighty := r.FormValue("toneighty")
		NineHtr := r.FormValue("ninehtr")
		MVP := r.FormValue("mvp")
		HighOut := r.FormValue("highout")
		SixBulls := r.FormValue("sixbulls")
		var Five_singles_win int
		var Five_singles_loss int
		if Five_singles_wldnp == "Win" {
			Five_singles_win = 1
			Five_singles_loss = 0
		} else if Five_singles_wldnp == "Loss" {
			Five_singles_win = 0
			Five_singles_loss = 1
		} else {
			Five_singles_win = 0
			Five_singles_loss = 0
		}
		var Cricket_singles_win int
		var Cricket_singles_loss int
		if Cricket_singles_wldnp == "Win" {
			Cricket_singles_win = 1
			Cricket_singles_loss = 0

		} else if Cricket_singles_wldnp == "Loss" {
			Cricket_singles_win = 0
			Cricket_singles_loss = 1
		} else {
			Cricket_singles_win = 0
			Cricket_singles_loss = 0
		}
		var Five_doubles_win int
		var Five_doubles_loss int
		if Five_doubles_wldnp == "Win" {
			Five_doubles_win = 1
			Five_doubles_loss = 0
		} else if Five_doubles_wldnp == "Loss" {
			Five_doubles_win = 0
			Five_doubles_loss = 1
		} else {
			Five_doubles_win = 0
			Five_doubles_loss = 0
		}
		var Cricket_doubles_win int
		var Cricket_doubles_loss int
		if Cricket_doubles_wldnp == "Win" {
			Cricket_doubles_win = 1
			Cricket_doubles_loss = 0
		} else if Cricket_doubles_wldnp == "Loss" {
			Cricket_doubles_win = 0
			Cricket_doubles_loss = 1
		} else {
			Cricket_doubles_win = 0
			Cricket_doubles_loss = 0
		}
		fmt.Println("INSERT Five_singles_dnp ", Five_singles_wldnp)
		insForm, err := db.Prepare("INSERT INTO Player_Stats(week,team,player,five_singles_wldnp, five_singles_win,five_singles_loss,cricket_singles_wldnp,cricket_singles_win,cricket_singles_loss,five_doubles_wldnp,five_doubles_win,five_doubles_loss,cricket_doubles_wldnp,cricket_doubles_win,cricket_doubles_loss,plusone,toneighty,ninehtr,mvp,highout,sixbulls) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
		if err != nil {
			fmt.Println("New insert insForm top")
			panic(err.Error())
		}
		insForm.Exec(Week, Team, Player, Five_singles_wldnp, Five_singles_win, Five_singles_loss, Cricket_singles_wldnp, Cricket_singles_win, Cricket_singles_loss, Five_doubles_wldnp, Five_doubles_win, Five_doubles_loss, Cricket_doubles_wldnp, Cricket_doubles_win, Cricket_doubles_loss, PlusOne, TonEighty, NineHtr, MVP, HighOut, SixBulls)
		log.Println("INSERT: Player: " + Player + " | Week: " + holdweek + " | Team: " + Team)
	}
	defer db.Close()
	posted = "N"
	http.Redirect(w, r, "/", 301)
}
