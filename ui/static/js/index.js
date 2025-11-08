const getStudentsBtn = document.getElementById('get-students-btn');
const addStudentBtn = document.getElementById('add-student-btn');
const addFirstNameInput = document.getElementById('add-first-name-input')
const addLastNameInput = document.getElementById('add-last-name-input')
const addEmailInput = document.getElementById('add-email-input')
const addDateInput = document.getElementById('add-date-input')

/**
 * @type {HTMLInputElement[]|null}
 */
const addInputList = [addFirstNameInput, addLastNameInput, addEmailInput, addDateInput];
/**
 * @type {Student[]}
 */
let students = [];