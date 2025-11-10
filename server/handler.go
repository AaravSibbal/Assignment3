package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AaravSibbal/COMP3005Assignment3/pkg/psql"
	"github.com/AaravSibbal/COMP3005Assignment3/pkg/student"
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
	studentList, err := psql.GetAllStudents(app.db, app.ctx);
	if err != nil {
		app.serverError(w, err)
		return
	}
	studentListJson, err := json.Marshal(studentList);
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(studentListJson)
}

func (app *application) addStudent(w http.ResponseWriter, r *http.Request){
	var student student.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		app.serverError(w, err);
		return
	}
	fmt.Printf("Recieved Student: %+v", student);
	err = psql.AddStudent(app.db, app.ctx, &student)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(psql.ConvertErrorToJsonObj(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(psql.SuccessMessage{Message: "Student was created"})
}

func (app *application) updateEmail(w http.ResponseWriter, r *http.Request){
	var student student.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		app.serverError(w, err)
		return
	}
	fmt.Printf("Recieved Student Data: %+v\n", student)
	err = psql.UpdateEmail(app.db, app.ctx, &student)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(psql.ConvertErrorToJsonObj(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(psql.SuccessMessage{Message: "Email was updated successfully"})
}

func (app *application) deleteStudent(w http.ResponseWriter, r *http.Request){
	var student student.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		app.serverError(w, err)
		return
	}
	fmt.Printf("Recieved Student Data: %+v\n", student)
	err = psql.DeleteStudent(app.db, app.ctx, &student)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(psql.ConvertErrorToJsonObj(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(psql.SuccessMessage{Message: "Student was deleted"})
}