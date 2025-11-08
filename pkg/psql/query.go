package psql

import (
	"context"
	"database/sql"
	// "fmt"
	"time"

	"github.com/AaravSibbal/COMP3005Assignment3/pkg/student"
)

func GetAllStudents(db *sql.DB, ctx *context.Context) (*student.StudentList, error) {
	newCtx, cancel := context.WithTimeout(*ctx, time.Second*5)
	defer cancel()

	stmt, err := db.PrepareContext(newCtx, "SELECT * FROM students;")
	if err != nil {
		return nil, err;
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(newCtx)
	if err != nil {
		return nil, err
	}

	studentList := student.CreateStudentList()
	for(rows.Next()){
		studentID := 0
		firstName := ""
		lastName := ""
		email := ""
		enrollmentDate := time.Time{}

		err = rows.Scan(&studentID, &firstName, &lastName, &email, &enrollmentDate)
		if err != nil{
			return nil, err
		}

		stu := student.CreateStudent(studentID, firstName, lastName, email, &enrollmentDate)
		studentList.AddStudent(stu)
	}
	studentList.Print()
	return studentList, nil
}