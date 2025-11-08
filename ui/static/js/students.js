class Student {
    studentId
    firstName
    lastName
    email
    enrollmentDate

    constructor(studentId, firstName, lastName, email, enrollmentDate){
        this.studentsId = studentId;
        this.firstName = firstName;
        this.lastName = lastName;
        this.email = email;
        this.enrollmentDate = enrollmentDate;
    }
}

function createStudentArrFromJson(obj){
    students = []
    
    let len = obj.length

    for(let i=0; i<len; i++){
        let studentObj = obj.students[i]
        let student = new Student(studentObj.students_id, 
            studentObj.first_name, studentObj.last_name, 
            studentObj.email, studentObj.enrollment_date);
            students.push(student)
        // let playerHTMLRow = player.toHTMLTableRow()
        // rankingTableBody.appendChild(playerHTMLRow)
    }
    console.log(students);
}
