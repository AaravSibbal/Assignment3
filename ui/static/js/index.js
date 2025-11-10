const getStudentsBtn = document.getElementById('get-students-btn');
const addStudentBtn = document.getElementById('add-student-btn');
const addFirstNameInput = document.getElementById('add-first-name-input')
const addLastNameInput = document.getElementById('add-last-name-input')
const addEmailInput = document.getElementById('add-email-input')
const addDateInput = document.getElementById('add-date-input')
const addStudentResponseError = document.getElementById('add-student-response-error')
const updateIdInput = document.getElementById('update-id-input')
const updateEmailInput = document.getElementById('update-email-input')
const updateEmailBtn = document.getElementById('update-email-btn')
const updateEmailResponseError = document.getElementById('update-email-response-error')
const deleteIdInput = document.getElementById('delete-id-input')
const deleteStudentBtn = document.getElementById('delete-student-btn')
const deleteStudentResponseError = document.getElementById('delete-student-response-error')

/**
 * @type {HTMLInputElement[]|null}
 */
const addInputList = [addFirstNameInput, addLastNameInput, addEmailInput, addDateInput];
const addErrorMsgId = ['add-first-name-error', 'add-last-name-error','add-date-error', 'add-email-error'] 

/**
 * @type {HTMLInputElement[]|null}
 */
const updateInputList = [updateEmailInput, updateIdInput];
const updateErrorMsgId = ['update-id-input', 'update-email-input']

const deleteInputList = [deleteIdInput]
const deleteErrorMsgId = ['delete-id-error']
/**
 * @type {Student[]}
 */
let students = [];