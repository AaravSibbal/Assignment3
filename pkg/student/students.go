package student

import (
	"encoding/json"
	"fmt"
	"time"
)

type Student struct {
	StudentsId     int        `json:"students_id"`
	FirstName      string     `json:"first_name"`
	LastName       string     `json:"last_name"`
	Email          string     `json:"email"`
	EnrollmentDate *time.Time `json:"enrollment_date"`
}

type StudentList struct {
	Lenght     int        `json:"length"`
	StudentArr []*Student `json:"students"`
}

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

func CreateStudentList() *StudentList {
	stuList := &StudentList{
		Lenght:     0,
		StudentArr: make([]*Student, 10),
	}

	return stuList
}

func (sl *StudentList) AddStudent(newStu *Student) error {
	if newStu == nil {
		return fmt.Errorf("passed nill pointer for student, it is not initialized")
	}
	if sl.Lenght == cap(sl.StudentArr) {
		sl.resize()
	}
	sl.StudentArr[sl.Lenght] = newStu
	sl.Lenght++
	return nil
}

func (sl *StudentList) resize() {
	newList := make([]*Student, 2*len(sl.StudentArr))
	copy(newList, sl.StudentArr)
	sl.StudentArr = newList
}

func (sl *StudentList) MarshalJSON() ([]byte, error) {
	output := StudentList{
		Lenght:     sl.Lenght,
		StudentArr: sl.StudentArr,
	}

	return json.Marshal(output)
}

func (sl *StudentList) Print(){
	for i, stu:= range(sl.StudentArr) {
		if(stu == nil){
			continue;
		}
		fmt.Printf("%d) %d, %s, %s, %s, %v\n", i, stu.StudentsId, stu.FirstName, stu.LastName, stu.Email, stu.EnrollmentDate)
	}
}