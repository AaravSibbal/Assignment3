class Student {
    student_id
    first_name
    last_name
    email
    enrollment_date

    constructor(studentId, firstName, lastName, email, enrollmentDate){
        this.student_id = studentId;
        this.first_name = firstName;
        this.last_name = lastName;
        this.email = email;
        this.enrollment_date = enrollmentDate;
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

function createStudentFromInput(){
    let firstName = addFirstNameInput.value
    let lastName = addLastNameInput.value
    let email = addEmailInput.value
    let date = `${addDateInput.value}T00:00:00Z`

    let student = new Student(null, firstName, lastName, email, date);
    return student
}

/**
 * clear out the error messages in the add student form
 */
function clearErrorMsgs(){
    for(let i=0; i<errorMsgId.length; i++){
        document.getElementById(errorMsgId[i]).innerText = ""
    }
    addStudentResponseError.innerText = ""
}

function addStudent(){
    clearErrorMsgs()
    // check for the empty field
    let isFieldEmpty = false;
    for(let i=0; i<addInputList.length; i++){
        if(addInputList[i].value == ""){
            let id = addInputList[i].id
            let errorId = id.substring(0, id.length-6) +"-error"
            console.log(errorId);
            let errorElem = document.getElementById(errorId)
            errorElem.innerText = "This field is required" 
            isFieldEmpty = true
        }
    }
    if(isFieldEmpty){
        return
    }
    let student = createStudentFromInput()
    console.log(student);
    addStudentInDB( student)
}
