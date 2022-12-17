function Register() {
    const username = document.getElementById("username").value
    const password = document.getElementById("password").value
    const email = document.getElementById("email").value
    fetch("../../register", {
        method: "POST",
        headers: { "Content-Type": "application/json"},
        body: JSON.stringify({
            username: username,
            password: password,
            email: email
        })
    }).then(res => {
        if (!res.ok) {
            throw new Error("response was not ok")
        }
        document.location.href = '/'
    })
}
