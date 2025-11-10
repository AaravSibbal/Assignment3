package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	// "fmt"
	"time"

	"github.com/AaravSibbal/COMP3005Assignment3/pkg/student"
	"github.com/lib/pq"
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

func AddStudent(db *sql.DB, ctx *context.Context, stu *student.Student) (error) {
	newCtx, cancel := context.WithTimeout(*ctx, time.Second*5)
	defer cancel();

	stmt, err := db.PrepareContext(newCtx, "INSERT INTO students (first_name, last_name, email, enrollment_date) VALUES ($1, $2, $3, $4);")
	if err != nil{
		return err
	}

	_, err = stmt.ExecContext(newCtx, stu.FirstName, stu.LastName, stu.Email, stu.EnrollmentDate)
	var pqErr *pq.Error
	
	if errors.As(err, &pqErr){
		switch PsqlErrors(pqErr.Code) {
		case UniqueConstraintError:
			fmt.Printf("DB ERROR: Unique constraint is broken on: %s\n", pqErr.Constraint)
			return errors.New("email already exists")
		case NotNullError:
			fmt.Printf("DB ERROR: not null voilation on column: %s\n", pqErr.Column)
			return fmt.Errorf("required field is missing: %s", pqErr.Column)
		default:
			fmt.Printf("Unhandeled psql state: %s - %s\n", pqErr.Code, pqErr.Message)
			return errors.New("an unexpected database error has occurred")
		}	
	}

	return nil
}

func UpdateEmail(db *sql.DB, ctx *context.Context, stu *student.Student) error {


	return nil
}