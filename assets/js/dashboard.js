function getMyCookies() {
    fetch("../../recipes", {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": "Bearer " + sessionStorage.getItem('token')
        }
    }).then(res => {
        if (!res.ok) {
            throw new Error("response was not ok")
        }
        res.json().then(data => {
            console.log(data)
            recipeList = document.getElementById("recipe-list")
            recipeList.innerHTML = generateRecipeList(data)
            console.log(recipeList.innerHTML)
        })
    })
}

function searchRecipes() {
    const filter = document.getElementById("recipe-search").value
    const recipeList = document.getElementById("recipe-list-search")
    if (filter.length >= 2) {
        fetch(`../../recipes/name-like/${filter}`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": "Bearer " + sessionStorage.getItem('token')
            }
        }).then(res => {
            if (!res.ok) {
                if (res.status == 404) {
                    recipeList.innerHTML = "no recipes found"
                    return
                } else {
                    recipeList.innerHTML = ""
                    throw new Error("response was not ok")
                }
            }
            res.json().then(data => {
                console.log(data)
                recipeList.innerHTML = generateRecipeList(data)
                console.log(recipeList.innerHTML)
            })
        })
    }

}

function generateRecipeList(data) {
    return `
    <ul class="recipe-list">
                    ${data.map(recipe =>
        `<li class="recipe-list-element" id="recipe-list-id-${recipe.id}">
                        ${recipe.name}
                        <ul> ${recipe.ingredients.map(ing => `<li class="recipe-list-ingredient">${ing.name} ${ing.amount} ${ing.unit}</li>`).join('')}</ul>
                        ${recipe.text}
                     </li>`
    ).join('')}
    </ul>`
}