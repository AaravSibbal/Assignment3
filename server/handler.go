package server

import (
	// "encoding/json"
	"net/http"

	// psql "github.com/AaravSibbal/COMP3005Assignment3/pkg/sql"
)

func (app *application) pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	htmlFile, err := app.readHTMLFile("index.html")
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(htmlFile)
}

func (app *application) getStudents(w http.ResponseWriter, r *http.Request) {
	
}

// func (app *application) playerRankings(w http.ResponseWriter, r *http.Request) {
// 	players, err := psql.GetRanking(app.db, app.ctx)
// 	if err != nil {
// 		app.serverError(w, err)
// 		return
// 	}

// 	playersJson, err := json.Marshal(players)
// 	if err != nil {
// 		app.serverError(w, err)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(playersJson)
// }

// // func (app *application) playerStat(w http.ResponseWriter, r *http.Request) {
// // 	name := r.URL.Query().Get(":name")
// // 	if name == "" {
// // 		app.clientError(w, http.StatusBadRequest)
// // 		return
// // 	}

// // 	player, err := psql.GetPlayerData(app.db, app.ctx, name)
// // 	if err != nil {
// // 		app.serverError(w, err)
// // 		return
// // 	}

// // 	playerJson, err := json.Marshal(player)
// // 	if err != nil {
// // 		app.serverError(w, err)
// // 		return
// // 	}

// // 	w.Header().Set("Content-Type", "application/json")
// // 	w.Write(playerJson)
// // }

// func (app *application) playerHtml(w http.ResponseWriter, r *http.Request) {
// 	html, err := app.readHTMLFile("player.html")
// 	if err != nil {
// 		app.serverError(w, err)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "text/html")
// 	w.Write(html)
// }
