function getStudents(){
    fetch("/students", {
        method: "GET",
    }).then((response)=>{
        if (!response.ok) {
            throw new Error("response was not okay")
        }

        return response.json()
    }).then(jsonObj=>{
        console.log(jsonObj)
        createStudentArrFromJson(jsonObj)
    }).catch(error => {
        console.error("there was a problem with the fetch operating\n\n"+error)
    })
}

function addStudentInDB(student){
    console.log(JSON.stringify(student));
    
    fetch("/student/add", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(student)
    })
    .then(async response=>{
        if(!response.ok){
            const errorBody = await response.json()
            addStudentResponseError.innerText = errorBody.message 
            throw new Error("http Error")
        }
        return response.json()
    })
    .then(jsonObj=>{
        console.log(jsonObj)
    })
    .catch(error=>{
        console.error("there was an error adding our boy: n\n\n"+error)
    })
}

function updateStudentEmailInDb(student){
    console.log(JSON.stringify(student));
    
    fetch("/student/email/update", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(student)
    })
    .then(async response=>{
        if(!response.ok){
            const errorBody = await response.json()
            updateEmailResponseError.innerText = errorBody.message 
            throw new Error("http Error")
        }
        return response.json()
    })
    .then(jsonObj=>{
        console.log(jsonObj)
    })
    .catch(error=>{
        console.error("there was an error adding our boy: n\n\n"+error)
    })
}

