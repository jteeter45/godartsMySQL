package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func Update(w http.ResponseWriter, r *http.Request) {
	posted = "N"
	db := dbConn()
	if r.Method == "POST" {
		fmt.Println("In Update ID is :", r.FormValue("id"))
		Id := r.FormValue("id")
		Week := r.FormValue("week")
		Team := r.FormValue("team")
		Player := r.FormValue("player")
		Five_singles_wldnp := r.FormValue("five_singles_wldnp")
		Cricket_singles_wldnp := r.FormValue("cricket_singles_wldnp")
		Five_doubles_wldnp := r.FormValue("five_doubles_wldnp")
		Cricket_doubles_wldnp := r.FormValue("cricket_doubles_wldnp")
		PlusOne := r.FormValue("plusone")
		TonEighty := r.FormValue("toneighty")
		NineHtr := r.FormValue("ninehtr")
		MVP := r.FormValue("mvp")
		HighOut := r.FormValue("highout")
		SixBulls := r.FormValue("sixbulls")
		var Five_singles_win string
		var Five_singles_loss string
		if Five_singles_wldnp == "Win" {
			Five_singles_win = "1"
			Five_singles_loss = "0"
		} else if Five_singles_wldnp == "Loss" {
			Five_singles_win = "0"
			Five_singles_loss = "1"
		} else {
			Five_singles_win = "0"
			Five_singles_loss = "0"
		}
		strconv.Atoi(Five_singles_win)
		strconv.Atoi(Five_singles_loss)
		var Cricket_singles_win string
		var Cricket_singles_loss string
		if Cricket_singles_wldnp == "Win" {
			Cricket_singles_win = "1"
			Cricket_singles_loss = "0"

		} else if Cricket_singles_wldnp == "Loss" {
			Cricket_singles_win = "0"
			Cricket_singles_loss = "1"
		} else {
			Cricket_singles_win = "0"
			Cricket_singles_loss = "0"
		}
		strconv.Atoi(Cricket_singles_win)
		strconv.Atoi(Cricket_singles_loss)
		var Five_doubles_win string
		var Five_doubles_loss string
		if Five_doubles_wldnp == "Win" {
			Five_doubles_win = "1"
			Five_doubles_loss = "0"
		} else if Five_doubles_wldnp == "Loss" {
			Five_doubles_win = "0"
			Five_doubles_loss = "1"
		} else {
			Five_doubles_win = "0"
			Five_doubles_loss = "0"
		}
		strconv.Atoi(Five_doubles_win)
		strconv.Atoi(Five_doubles_loss)
		var Cricket_doubles_win string
		var Cricket_doubles_loss string
		if Cricket_doubles_wldnp == "Win" {
			Cricket_doubles_win = "1"
			Cricket_doubles_loss = "0"
		} else if Cricket_doubles_wldnp == "Loss" {
			Cricket_doubles_win = "0"
			Cricket_doubles_loss = "1"
		} else {
			Cricket_doubles_win = "0"
			Cricket_doubles_loss = "0"
		}
		strconv.Atoi(Cricket_doubles_win)
		strconv.Atoi(Cricket_doubles_loss)
		insForm, err := db.Prepare("UPDATE Player_Stats SET Week=?, Team=?, Player=?, Five_singles_wldnp=?,Five_singles_win=?,Five_singles_loss=?,Cricket_singles_wldnp=?,Cricket_singles_win=?,Cricket_singles_loss=?,Five_doubles_wldnp=?,Five_doubles_win=?,Five_doubles_loss=?,Cricket_doubles_wldnp=?,Cricket_doubles_win=?,Cricket_doubles_loss=?,PlusOne=?,TonEighty=?,NineHtr=?,MVP=?,HighOut=?,SixBulls=? WHERE Id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(Week, Team, Player, Five_singles_wldnp, Five_singles_win, Five_singles_loss, Cricket_singles_wldnp, Cricket_singles_win, Cricket_singles_loss, Five_doubles_wldnp, Five_doubles_win, Five_doubles_loss, Cricket_doubles_wldnp, Cricket_doubles_win, Cricket_doubles_loss, PlusOne, TonEighty, NineHtr, MVP, HighOut, SixBulls, Id)
		log.Println("UPDATE Complete")

		fmt.Println("Endpoint Hit: Update")
		log.Println("UPDATE: Player: " + Player + " | Week: " + Week + " | Team: " + Team)
		var path = "test.json"
		var _ = os.Remove(path)
		//if err != nil {
		//panic(err.Error())
		//}
	}
	posted = "N"
	http.Redirect(w, r, "/", 301)
}
