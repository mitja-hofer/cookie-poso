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

function searchRecipesByName() {
    const filter = document.getElementById("recipe-search").value
    const recipeList = document.getElementById("recipe-list-search")
    recipeList.innerHTML = ""
    fetch(`../../recipes/name/${filter}`, {
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
                throw new Error("response was not ok")
            }
        }
        res.json().then(data => {
            recipeList.innerHTML = generateRecipeList(data)
        })
    })
}

function searchRecipesByIngredient() {
    const filter = document.getElementById("recipe-search").value
    const recipeList = document.getElementById("recipe-list-search")
    recipeList.innerHTML = ""
    fetch(`../../recipes/ingredient/${filter}`, {
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
                throw new Error("response was not ok")
            }
        }
        res.json().then(data => {
            recipeList.innerHTML = generateRecipeList(data)
        })
    })
}

function searchIngredients() {
    const filter = document.getElementById("ingredient-search").value
    const ingredientList = document.getElementById("ingredient-list-search")
    if (filter.length >= 2) {
        fetch(`../../ingredients/name-like/${filter}`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": "Bearer " + sessionStorage.getItem('token')
            }
        }).then(res => {
            if (!res.ok) {
                if (res.status == 404) {
                    ingredientList.innerHTML = "no ingredients found"
                    return
                } else {
                    ingredientList.innerHTML = ""
                    throw new Error("response was not ok")
                }
            }
            res.json().then(data => {
                ingredientList.innerHTML = generateIngredientList(data)
            })
        })
    } else {
        ingredientList.innerHTML = ""
    }
}


function generateRecipeList(data) {
    if (data == null) {
        return "No recipes found."
    }
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

function generateIngredientList(data) {
    if (data == null) {
        return "No ingredients found."
    }
    return `
    <ul class="ingredient-list">
                    ${data.map(ingredient =>
        `<li class="ingredient-list-element" id="ingredient-list-id-${ingredient.id}">
                        ${ingredient.name}
                     </li>`
    ).join('')}
    </ul>`

}