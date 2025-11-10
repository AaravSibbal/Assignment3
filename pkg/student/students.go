package student

import (
	"encoding/json"
	"fmt"
	"time"
)

// student struct with the respective json info
type Student struct {
	StudentsId     int        `json:"student_id,omitempty"`
	FirstName      string     `json:"first_name"`
	LastName       string     `json:"last_name"`
	Email          string     `json:"email"`
	EnrollmentDate *time.Time `json:"enrollment_date,omitempty"`
}

// student list struct to make adding students easier
type StudentList struct {
	Lenght     int        `json:"length"`
	StudentArr []*Student `json:"students"`
}

// function foe craeting students
func CreateStudent(studentId int, firstName string, lastName string,
	email string, enrollmentDate *time.Time) *Student {
	stu := &Student{
		StudentsId:     studentId,
		FirstName:      firstName,
		LastName:       lastName,
		Email:          email,
		EnrollmentDate: enrollmentDate,
	}

	return stu
}

// function for creating the studentList
func CreateStudentList() *StudentList {
	stuList := &StudentList{
		Lenght:     0,
		StudentArr: make([]*Student, 10),
	}

	return stuList
}

// function for adding the student
func (sl *StudentList) AddStudent(newStu *Student) error {
	// check for null pointers
	if newStu == nil {
		return fmt.Errorf("passed nill pointer for student, it is not initialized")
	}
	// check if the backing array is full
	if sl.Lenght == cap(sl.StudentArr) {
		sl.resize()
	}
	// add the student at the back of the list and increase the size
	sl.StudentArr[sl.Lenght] = newStu
	sl.Lenght++
	return nil
}

// resize function makes the array 2 times the size of the elems
func (sl *StudentList) resize() {
	newList := make([]*Student, 2*len(sl.StudentArr))
	copy(newList, sl.StudentArr)
	sl.StudentArr = newList
}

// for converting the studentList into json (custom function)
func (sl *StudentList) MarshalJSON() ([]byte, error) {
	output := StudentList{
		Lenght:     sl.Lenght,
		StudentArr: sl.StudentArr,
	}

	return json.Marshal(output)
}

// for debugging a print function
func (sl *StudentList) Print(){
	for i, stu:= range(sl.StudentArr) {
		if(stu == nil){
			continue;
		}
		fmt.Printf("%d) %d, %s, %s, %s, %v\n", i, stu.StudentsId, stu.FirstName, stu.LastName, stu.Email, stu.EnrollmentDate)
	}
}