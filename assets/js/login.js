function postLoginForm() {
    const username = document.getElementById("username").value
    const password = document.getElementById("password").value
    fetch("login", {
        method: "POST",
        headers: { "Content-Type": "application/json"},
        body: JSON.stringify({
            username: username,
            password: password
        })
    }).then(res => {
        if (!res.ok) {
            throw new Error("response ws not ok")
        }
        res.json()
            .then(data => {
                console.log(data)
                sessionStorage.setItem("token", data.token)
                document.location.href = '/assets/html/dashboard'
            })
    })
}