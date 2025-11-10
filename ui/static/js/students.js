
// student class
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

// converts the jsonObj to a Students Arr
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
 * adds a html cell to the row provided
 * @param {HTMLTableRowElement} row 
 * @param {*} value 
 */
function addCellToRow(row, value){
    const cell = row.insertCell()
    cell.textContent = value
}

/**
 * converts the student arr to a html table
 * @param {Student[]} studentArr 
 */
function studentArrToTable(studentArr){
    // make sure the table body is empty again
    tableBody.innerHTML = ""
    // populate the table body with fresh data
    for(let i=0; i<studentArr.length; i++){
        const stu = studentArr[i]
        // create the row
        const row = tableBody.insertRow();
        // populate the row
        addCellToRow(row, stu.student_id)
        addCellToRow(row, stu.first_name)
        addCellToRow(row, stu.last_name)
        addCellToRow(row, stu.email)
        addCellToRow(row, stu.enrollment_date)
    }
}

/**
 * takes the input from the add students 
 * params and makes a student object from it
 * @returns {Student}
 */
function createStudentFromInput(){
    let firstName = addFirstNameInput.value
    let lastName = addLastNameInput.value
    let email = addEmailInput.value
    // this is some golang shenanigans
    let date = `${addDateInput.value}T00:00:00Z`

    let student = new Student(null, firstName, lastName, email, date);
    return student
}

/**
 * takes in the ids as an array of string for the functions
 * add, delete, update
 * and the response error element for all those respective things
 * clears them out so they can be repopulated appropriately
 * @param {string[]} errorIds 
 * @param {HTMLParagraphElement} responseErrorElem
 */
function clearErrorMsgs(errorIds, responseErrorElem){
    for(let i=0; i<errorIds.length; i++){
        document.getElementById(errorIds[i]).innerText = ""
    }
    responseErrorElem.innerText = ""
}

/**
 * takes in an input elements list and makes sure they are
 * filled shows a error msg if they are not
 * @param {HTMLInputElement[]} inputElemList 
 * @returns {void}
 */
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

/**
 * add student the client side
 * @returns {void}
 */
function addStudent(){
    // clear the error messages
    clearErrorMsgs(addErrorMsgId, addStudentResponseError)
    // check for the empty field if it is return early
    let isFieldEmpty = areInputFieldsFilled(addInputList)
    if(isFieldEmpty){
        return
    }
    // create the student from input
    let student = createStudentFromInput()
    console.log(student);
    // call the server to add the student
    addStudentInDB(student)
}

function updateStudentEmail(){
    // clear the error messages
    clearErrorMsgs(updateErrorMsgId, updateEmailResponseError)
    // make sure the input field are filled
    let isFieldEmpty = areInputFieldsFilled(updateInputList)
    if(isFieldEmpty){
        return
    }

    // make sure the ID field is only taking in an integer
    if(parseInt(updateIdInput.value) != parseFloat(updateIdInput.value)){
        document.getElementById('update-id-error').innerText = "This value has to be an integer"
        return
    }
    // create a student object
    let student = new Student(parseInt(updateIdInput.value), null, null, updateEmailInput.value, null)
    console.log(student)
    // call the server to update that student
    updateStudentEmailInDb(student)
}

function deleteStudent(){
    // clear the error messages
    clearErrorMsgs(deleteErrorMsgId, deleteStudentResponseError)
    // check to see if the input field are filled
    let isFieldEmpty = areInputFieldsFilled(deleteInputList)
    if(isFieldEmpty){
        return
    }
    // check if the id field is an integer
    if(parseInt(updateIdInput.value) != parseFloat(updateIdInput.value)){
        document.getElementById('update-id-error').innerText = "This value has to be an integer"
        return
    }
    let studentId = parseInt(deleteIdInput.value)
    // create the student object
    let student = new Student(studentId, null, null, null, null)
    console.log(student)
    // call the server to delete that student
    deleteStudentInDb(student)
}
