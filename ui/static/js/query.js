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