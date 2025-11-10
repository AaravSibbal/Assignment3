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
        let student = new Student(studentObj.student_id, 
            studentObj.first_name, studentObj.last_name, 
            studentObj.email, studentObj.enrollment_date);
            students.push(student)
        // let playerHTMLRow = player.toHTMLTableRow()
        // rankingTableBody.appendChild(playerHTMLRow)
    }
    console.log(students);
    return students;
}

/**
 * 
 * @param {HTMLTableRowElement} row 
 * @param {*} value 
 */
function addCellToRow(row, value){
    const cell = row.insertCell()
    cell.textContent = value
}

/**
 * 
 * @param {Student[]} studentArr 
 */
function studentArrToTable(studentArr){
    tableBody.innerHTML = ""
    for(let i=0; i<studentArr.length; i++){
        const stu = studentArr[i]
        const row = tableBody.insertRow();
        addCellToRow(row, stu.student_id)
        addCellToRow(row, stu.first_name)
        addCellToRow(row, stu.last_name)
        addCellToRow(row, stu.email)
        addCellToRow(row, stu.enrollment_date)
    }
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

/**
 * 
 * @param {string[]} errorIds 
 * @param {HTMLParagraphElement} responseErrorElem
 */
function clearErrorMsgs(errorIds, responseErrorElem){
    for(let i=0; i<errorIds.length; i++){
        document.getElementById(errorIds[i]).innerText = ""
    }
    responseErrorElem.innerText = ""
}

function areInputFieldsFilled(inputElemList){
    let isFieldEmpty = false
    for(let i=0; i<inputElemList.length; i++){
        if(inputElemList[i].value == "" || inputElemList[i].value == null){
            let id = inputElemList[i].id
            let errorId = id.substring(0, id.length-6)+"-error"
            let errorElem = document.getElementById(errorId)
            errorElem.innerText = "This field is required"
            isFieldEmpty = true
        }
    }
    return isFieldEmpty

}

function addStudent(){
    clearErrorMsgs(addErrorMsgId, addStudentResponseError)
    // check for the empty field
    let isFieldEmpty = areInputFieldsFilled(addInputList)
    if(isFieldEmpty){
        return
    }
    let student = createStudentFromInput()
    console.log(student);
    addStudentInDB(student)
}

function updateStudentEmail(){
    clearErrorMsgs(updateErrorMsgId, updateEmailResponseError)
    let isFieldEmpty = areInputFieldsFilled(updateInputList)
    if(isFieldEmpty){
        return
    }

    if(parseInt(updateIdInput.value) != parseFloat(updateIdInput.value)){
        document.getElementById('update-id-error').innerText = "This value has to be an integer"
        return
    }
    let student = new Student(parseInt(updateIdInput.value), null, null, updateEmailInput.value, null)
    console.log(student)
    updateStudentEmailInDb(student)
}

function deleteStudent(){
    clearErrorMsgs(deleteErrorMsgId, deleteStudentResponseError)
    let isFieldEmpty = areInputFieldsFilled(deleteInputList)
    if(isFieldEmpty){
        return
    }
    if(parseInt(updateIdInput.value) != parseFloat(updateIdInput.value)){
        document.getElementById('update-id-error').innerText = "This value has to be an integer"
        return
    }
    let studentId = parseInt(deleteIdInput.value)
    let student = new Student(studentId, null, null, null, null)
    console.log(student)
    deleteStudentInDb(student)
}
