function uploadImage() {
    const form = document.getElementById("imageUploadForm")
    const input = document.querySelector("input[type=file]")
    const img = document.getElementById("img")
    const data = new FormData(form)
    data.append("upload", input.files[0])
    fetch("../../image", {
        method: "POST",
        headers: {
            "Authorization": "Bearer " + sessionStorage.getItem('token')
        },
        body: data
    }).then(res => {
        res.json().then(data => {
            console.log(data.url)
            img.src = data.url
        })
    })
}