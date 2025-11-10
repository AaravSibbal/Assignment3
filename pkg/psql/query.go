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

// calls the Db for getting all the students
func GetAllStudents(db *sql.DB, ctx *context.Context) (*student.StudentList, error) {
	// creates a times for 5 seconds for the query to execute
	newCtx, cancel := context.WithTimeout(*ctx, time.Second*5)
	defer cancel()

	/*
		prepares the statement so it can't be changed later
		by attackers
	*/
	stmt, err := db.PrepareContext(newCtx, "SELECT * FROM students;")
	if err != nil {
		return nil, err;
	}
	defer stmt.Close()

	// query the Db
	rows, err := stmt.QueryContext(newCtx)
	// handle errors
	if err != nil {
		return nil, err
	}

	// populate Student List
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
	// success
	return studentList, nil
}

func AddStudent(db *sql.DB, ctx *context.Context, stu *student.Student) (error) {
	// creates a times for 5 seconds for the query to execute
	newCtx, cancel := context.WithTimeout(*ctx, time.Second*5)
	defer cancel();

	/*
		prepares the statement so it can't be changed later
		by attackers
	*/
	stmt, err := db.PrepareContext(newCtx, "INSERT INTO students (first_name, last_name, email, enrollment_date) VALUES ($1, $2, $3, $4);")
	if err != nil{
		return err
	}

	// query the DB
	_, err = stmt.ExecContext(newCtx, stu.FirstName, stu.LastName, stu.Email, stu.EnrollmentDate)
	// handle errors
	if err != nil {
		var pqErr *pq.Error
		
		if errors.As(err, &pqErr){
			switch PsqlErrors(pqErr.Code) {
			// in case of a unique constraint violation
			case UniqueConstraintError:
				fmt.Printf("DB ERROR: Unique constraint is broken on: %s\n", pqErr.Constraint)
				return errors.New("email already exists")
			// in case of a not null constraint violation
			case NotNullError:
				fmt.Printf("DB ERROR: not null voilation on column: %s\n", pqErr.Column)
				return fmt.Errorf("required field is missing: %s", pqErr.Column)
			default:
				fmt.Printf("Unhandeled psql state: %s - %s\n", pqErr.Code, pqErr.Message)
				return errors.New("an unexpected database error has occurred")
			}	
		}
	}	
	// success
	return nil
}

func UpdateEmail(db *sql.DB, ctx *context.Context, stu *student.Student) error {
	// creates a times for 5 seconds for the query to execute
	newCtx, cancel := context.WithTimeout(*ctx, time.Second*5)
	defer cancel()
	
	/*
		prepares the statement so it can't be changed later
		by attackers
	*/
	stmt, err := db.PrepareContext(newCtx, "UPDATE students SET email = $1 WHERE student_id = $2;")
	if err != nil {
		return err
	}
	defer stmt.Close();

	_, err = stmt.ExecContext(newCtx, stu.Email, stu.StudentsId)
	if err != nil {
		var pqErr *pq.Error
		
		if errors.As(err, &pqErr){
			switch PsqlErrors(pqErr.Code) {
			// in case of a unique constraint violation
			case UniqueConstraintError:
				fmt.Printf("DB ERROR: Unique constraint is broken on: %s\n", pqErr.Constraint)
				return errors.New("email already exists")
			// in case of a not null constraint violation
			case NotNullError:
				fmt.Printf("DB ERROR: not null voilation on column: %s\n", pqErr.Column)
				return fmt.Errorf("required field is missing: %s", pqErr.Column)
			default:
				fmt.Printf("Unhandeled psql state: %s - %s\n", pqErr.Code, pqErr.Message)
				return errors.New("an unexpected database error has occurred")
			}	
		}
	}	
	// success
	return nil
}

func DeleteStudent(db *sql.DB, ctx *context.Context, stu *student.Student) error {
	// creates a times for 5 seconds for the query to execute
	newCtx, cancel := context.WithTimeout(*ctx, time.Second*5)
	defer cancel()
	
	/*
		prepares the statement so it can't be changed later
		by attackers
	*/
	stmt, err := db.PrepareContext(newCtx, "DELETE FROM students WHERE student_id = $1;")
	if err != nil {
		return err
	}
	defer stmt.Close();

	_, err = stmt.ExecContext(newCtx, stu.StudentsId)
	if err != nil {
		var pqErr *pq.Error
		
		if errors.As(err, &pqErr){
			switch PsqlErrors(pqErr.Code) {
			case UniqueConstraintError:
			// in case of a unique constraint violation
				fmt.Printf("DB ERROR: Unique constraint is broken on: %s\n", pqErr.Constraint)
				return errors.New("email already exists")
			// in case of a not null constraint violation
			case NotNullError:
				fmt.Printf("DB ERROR: not null voilation on column: %s\n", pqErr.Column)
				return fmt.Errorf("required field is missing: %s", pqErr.Column)
			default:
				fmt.Printf("Unhandeled psql state: %s - %s\n", pqErr.Code, pqErr.Message)
				return errors.New("an unexpected database error has occurred")
			}	
		}
	}	
	// success
	return nil
}