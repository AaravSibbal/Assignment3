package server

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

/*
routes file contains all the routes to the website 
and the functions that correspond to them
*/
func (app *application) Routes() http.Handler {

	standardMiddleware := alice.New(app.logRequest, app.recoverPanic, app.secureHeaders)

	mux := pat.New()

	mux.Get("/ping", standardMiddleware.ThenFunc(app.pong))
	mux.Get("/", standardMiddleware.ThenFunc(app.home))
	mux.Get("/students", standardMiddleware.ThenFunc(app.getStudents))
	mux.Post("/student/add", standardMiddleware.ThenFunc(app.addStudent))
	mux.Post("/student/email/update", standardMiddleware.ThenFunc(app.updateEmail))
	mux.Del("/student", standardMiddleware.ThenFunc(app.deleteStudent))

	fileServer := http.FileServer(http.Dir("ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)

}