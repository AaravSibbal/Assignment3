/**
 * this file has all the functions that call the server
 * for CRUD things
 */


function getStudents(){
    // call to server
    fetch("/students", {
        method: "GET",
    }).then((response)=>{
        // handle the response and error
        if (!response.ok) {
            throw new Error("response was not okay")
        }

        return response.json()
    }).then(jsonObj=>{
        // if there are no errors convert the json Obj to table
        console.log(jsonObj)
        // convert json Obj to Arr
        const stuArr = createStudentArrFromJson(jsonObj)
        // convert the Arr to table
        studentArrToTable(stuArr)

    }).catch(error => {
        // catch and handle errors
        console.error("there was a problem with the fetch operating\n\n"+error)
    })
}

function addStudentInDB(student){
    // post request to the sever
    fetch("/student/add", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(student)
    })
    .then(async response=>{
        // handle the response and the errors
        if(!response.ok){
            /**
             * convert the body to json
             * show the error msg to client
             * throw the error to break the chain
             */
            const errorBody = await response.json()
            addStudentResponseError.innerText = errorBody.message 
            throw new Error("http Error")
        }
        return response.json()
    })
    .then(jsonObj=>{
        // if success show the success msg
        console.log(jsonObj)
        addStudentResponseError.innerText = jsonObj.message
    })
    .catch(error=>{
        // handle the error
        console.error("there was an error adding our boy: n\n\n"+error)
    })
    .finally(()=>{
        // update the student list to show the change
        getStudents()
    })
}

function updateStudentEmailInDb(student){
    // call the server to update the email
    fetch("/student/email/update", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(student)
    })
    .then(async response=>{
        // handle the response and the errors
        if(!response.ok){
            // if there is an error 
            // read the body
            // show the error msg to client
            // throw the error to break the chain
            const errorBody = await response.json()
            updateEmailResponseError.innerText = errorBody.message 
            throw new Error("http Error")
        }
        return response.json()
    })
    .then(jsonObj=>{
        // if success show the success msg to the client
        console.log(jsonObj)
        updateEmailResponseError.innerText = jsonObj.message
    })
    .catch(error=>{
        // handle the error
        console.error("there was an error adding our boy: n\n\n"+error)
    })
    .finally(()=>{
        // update the students table to show the change in the db
        getStudents()
    })
}

function deleteStudentInDb(student){
    // send the request to the server to delete the student
    fetch("/student", {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(student)
    })
    .then(async response=>{
        // handle the reponse and the error
        if(!response.ok){
            // if there is an error 
            // read the body
            // show the error msg to client
            // throw the error to break the chain
            const errorBody = await response.json()
            updateEmailResponseError.innerText = errorBody.message 
            throw new Error("http Error")
        }
        return response.json()
    })
    .then(jsonObj=>{
        // if success show the msg to the client
        console.log(jsonObj)
        deleteStudentResponseError.innerText = jsonObj.message
    })
    .catch(error=>{
        // handle the error
        console.error("there was an error adding our boy: n\n\n"+error)
    })
    .finally(()=>{
        // update the table to show the change in the Db
        getStudents()
    })
}
