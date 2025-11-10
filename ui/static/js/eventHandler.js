/**
 * this file has all the event handlers
 */

// gets students
getStudentsBtn.addEventListener("click", ()=>{
    getStudents()
})

// add students
addStudentBtn.addEventListener('click', ()=>{
    addStudent()
})

// update students
updateEmailBtn.addEventListener('click', ()=>{
    updateStudentEmail()
})

// delete students
deleteStudentBtn.addEventListener('click', ()=>{
    deleteStudent()
})