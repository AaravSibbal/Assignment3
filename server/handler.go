package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AaravSibbal/COMP3005Assignment3/pkg/psql"
	"github.com/AaravSibbal/COMP3005Assignment3/pkg/student"
)

/*
	pings the server so see if it is alive
*/
func (app *application) pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

/*
	sends the index.html file
*/
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	htmlFile, err := app.readHTMLFile("index.html")
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write(htmlFile)
}

/*
	sends the student list stored in the db
*/
func (app *application) getStudents(w http.ResponseWriter, r *http.Request) {
	// make the db request
	studentList, err := psql.GetAllStudents(app.db, app.ctx);
	// handle errors
	if err != nil {
		app.serverError(w, err)
		return
	}
	// convert the result to json
	studentListJson, err := json.Marshal(studentList);
	if err != nil {
		app.serverError(w, err)
		return
	}

	// send the result
	w.Header().Set("Content-Type", "application/json")
	w.Write(studentListJson)
}

// adding the student in the Db
func (app *application) addStudent(w http.ResponseWriter, r *http.Request){
	// get the information in out student struct
	var student student.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		app.serverError(w, err);
		return
	}
	fmt.Printf("Recieved Student: %+v", student);
	
	// make the call to the DB
	err = psql.AddStudent(app.db, app.ctx, &student)

	w.Header().Set("Content-Type", "application/json")
	// handle the errors
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(psql.ConvertErrorToJsonObj(err))
		return
	}
	// successs case sending the json
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(psql.SuccessMessage{Message: "Student was created"})
}

// updating the student email
func (app *application) updateEmail(w http.ResponseWriter, r *http.Request){
	// get the student information in the student struct
	var student student.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	// handle errors
	if err != nil {
		app.serverError(w, err)
		return
	}
	fmt.Printf("Recieved Student Data: %+v\n", student)
	// call to the Db
	err = psql.UpdateEmail(app.db, app.ctx, &student)
	w.Header().Set("Content-Type", "application/json")
	// handle DB errors 
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(psql.ConvertErrorToJsonObj(err))
		return
	}
	// success case sending the json
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(psql.SuccessMessage{Message: "Email was updated successfully"})
}

// deleting the student in the Db
func (app *application) deleteStudent(w http.ResponseWriter, r *http.Request){
	// load the student information in the student struct
	var student student.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	// handle the errors
	if err != nil {
		app.serverError(w, err)
		return
	}
	fmt.Printf("Recieved Student Data: %+v\n", student)
	// call to the Db
	err = psql.DeleteStudent(app.db, app.ctx, &student)
	w.Header().Set("Content-Type", "application/json")
	// handle the Db error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(psql.ConvertErrorToJsonObj(err))
		return
	}
	// success case sending the json
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(psql.SuccessMessage{Message: "Student was deleted"})
}