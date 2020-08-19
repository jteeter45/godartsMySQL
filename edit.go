package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	fmt.Println("In Edit ID is :", nId)
	selDB, err := db.Query("SELECT * FROM Player_Stats WHERE id=?", nId)
	//selDB, err := db.Query("SELECT * FROM Player_Stats WHERE Player =? AND Week=?", nplayer, nweek)
	if err != nil {
		panic(err.Error())
	}
	ply := Player_Stats{}
	for selDB.Next() {
		var id, week, five_singles_win, five_singles_loss, cricket_singles_win, cricket_singles_loss, five_doubles_win, five_doubles_loss, cricket_doubles_win, cricket_doubles_loss, plusone, toneighty, ninehtr, mvp, highout, sixbulls int
		var team, player, five_singles_wldnp, cricket_singles_wldnp, five_doubles_wldnp, cricket_doubles_wldnp string
		err = selDB.Scan(&id, &week, &team, &player, &five_singles_wldnp, &five_singles_win, &five_singles_loss, &cricket_singles_wldnp, &cricket_singles_win, &cricket_singles_loss, &five_doubles_wldnp, &five_doubles_win, &five_doubles_loss, &cricket_doubles_wldnp, &cricket_doubles_win, &cricket_doubles_loss, &plusone, &toneighty, &ninehtr, &mvp, &highout, &sixbulls)
		if err != nil {
			panic(err.Error())
		}
		ply.Id = id
		ply.Player = player
		ply.Week = week
		ply.Team = team
		ply.Five_singles_wldnp = five_singles_wldnp
		ply.Cricket_singles_wldnp = cricket_singles_wldnp
		ply.Five_doubles_wldnp = five_doubles_wldnp
		ply.Cricket_doubles_wldnp = cricket_doubles_wldnp
		ply.PlusOne = plusone
		ply.TonEighty = toneighty
		ply.NineHtr = ninehtr
		ply.MVP = mvp
		ply.HighOut = highout
		ply.SixBulls = sixbulls
	}
	defer db.Close()
	tmpl.ExecuteTemplate(w, "Edit", ply)
}
