function listOfElementsByClassName(name) {
    return Array.from(document.getElementsByClassName(name));
}

function switchTab(event, tab) {
    listOfElementsByClassName("tab-content").forEach(item => item.style.display = "none")
    listOfElementsByClassName("tab-link").forEach(item => item.className = item.className.replace(" active", ""))
    document.getElementById(tab).style.display = "block"
    event.currentTarget.className += " active"
}