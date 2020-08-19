package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Index")
	fmt.Println("holdplayer: ", holdplayer)
	fmt.Println("holdteam: ", holdteam)
	fmt.Println("Holdweeknum: ", holdweeknum)
	fmt.Println("Posted: ", posted)
	if holdweeknum == 0 && holdteam == "" && holdplayer == "" && posted == "N" {
		fmt.Println("Endpoint Index Bare select")
		db := dbConn()
		selDB, err := db.Query("SELECT Id, Week, Team, Player FROM Player_Stats ORDER BY Week,Team,Player")
		if err != nil {
			panic(err.Error())
		}

		ply := Player_Stats{}
		res := []Player_Stats{}
		for selDB.Next() {
			var player, team string
			var id, week int
			err = selDB.Scan(&id, &week, &team, &player)
			if err != nil {
				panic(err.Error())
			}
			ply.Id = id
			ply.Player = player
			ply.Week = week
			ply.Team = team
			//fmt.Println(ply.Player)

			res = append(res, ply)
			//fmt.Println(res)
		}
		posted = "Y"
		defer db.Close()
		tmpl.ExecuteTemplate(w, "Index", res)
	}
	if holdweeknum != 0 && holdteam == "None" && posted == "N" {
		db := dbConn()
		fmt.Println("Index,Team is None, Select Week")
		selDB, err := db.Query("SELECT Id, Week, Team, Player  FROM Player_Stats WHERE Week=?", holdweeknum)
		if err != nil {
			panic(err.Error())
		}

		ply := Player_Stats{}
		res := []Player_Stats{}
		for selDB.Next() {
			var player, team string
			var id, week int
			err = selDB.Scan(&id, &week, &team, &player)
			if err != nil {
				panic(err.Error())
			}
			ply.Id = id
			ply.Player = player
			ply.Week = week
			ply.Team = team
			//fmt.Println(ply.Player)

			res = append(res, ply)

			//fmt.Println(res)
		}
		posted = "Y"
		defer db.Close()
		tmpl.ExecuteTemplate(w, "Index", res)
	}
	if holdweeknum != 0 && holdteam != "" && holdplayer != "" && posted == "N" {
		db := dbConn()
		fmt.Println("Index,Week,Team & Player all have values Select Week,Team,Player")
		selDB, err := db.Query("SELECT Id, Week, Team, Player FROM Player_Stats WHERE Week=? && Team=? && Player=?", holdweeknum, holdteam, holdplayer)
		if err != nil {
			panic(err.Error())
		}

		ply := Player_Stats{}
		res := []Player_Stats{}
		for selDB.Next() {
			var player, team string
			var id, week int
			err = selDB.Scan(&id, &week, &team, &player)
			if err != nil {
				panic(err.Error())
			}
			ply.Id = id
			ply.Player = player
			ply.Week = week
			ply.Team = team
			//fmt.Println(ply.Player)

			res = append(res, ply)

			fmt.Println(res)
		}
		posted = "Y"
		defer db.Close()
		tmpl.ExecuteTemplate(w, "Index", res)
	}
}
