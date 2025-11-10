package server

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) Routes() http.Handler {

	standardMiddleware := alice.New(app.logRequest, app.recoverPanic, app.secureHeaders)

	mux := pat.New()

	mux.Get("/ping", standardMiddleware.ThenFunc(app.pong))
	mux.Get("/", standardMiddleware.ThenFunc(app.home))
	mux.Get("/students", standardMiddleware.ThenFunc(app.getStudents))
	mux.Post("/student/add", standardMiddleware.ThenFunc(app.addStudent))
	mux.Post("/student/email/update", standardMiddleware.ThenFunc(app.updateEmail))
	// mux.Get("/player/ranking", standardMiddleware.ThenFunc(app.playerRankings))
	// mux.Get("/player/:name", standardMiddleware.ThenFunc(app.playerHtml))
	// mux.Get("/player/:name/stat", standardMiddleware.ThenFunc(app.playerStat))

	fileServer := http.FileServer(http.Dir("ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)

}