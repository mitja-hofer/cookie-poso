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
            recipeList = document.getElementById("recipe-list")
            recipeList.innerHTML = generateRecipeList(data)
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
            if (res.status === 404) {
                recipeList.innerHTML = "no recipes found"
                return
            } else {
                throw new Error("response was not ok")
            }
        }
        res.json().then(data => {
            recipeList.innerHTML = generateRecipeList([data])
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
    <ul class="recipe-list-ul">
                    ${data.map(recipe =>
        `<li class="recipe-list-element" id="recipe-list-id-${recipe.id}">
                        ${recipe.name}
                        <ul> ${recipe.ingredients ? recipe.ingredients.map(ing => `<li class="recipe-list-ingredient">${ing.name} ${ing.amount} ${ing.unit}</li>`).join(''): ""}</ul>
                        ${recipe.text}
                     </li>`
    ).join('')}
    </ul>`
}

function generateIngredientList(data) {
    const search = document.getElementById("ingredient-search").value
    if (data == null) {
        return `
            <ul class="ingredient-list">
                <li class="ingredient-list-element" id="ingredient-list-id-0">
                    New: ${search}
                    <input type="checkbox" id="ingredient-list-checkbox-id-0" value="0" data-name="${search}" onchange="handleTick(this)">
                    <div id="ingredient-list-details-id-0"></div>
                    </li>
            </ul>`
    }
    return `
    <ul class="ingredient-list">
                    ${data.map(ingredient =>
        `<li class="ingredient-list-element" id="ingredient-list-id-${ingredient.id}">
                        ${ingredient.name}
                        <input type="checkbox" id="ingredient-list-checkbox-id-${ingredient.id}" value="${ingredient.id}" data-name="${ingredient.name}" onchange="handleTick(this)">
                        <div id="ingredient-list-details-id-${ingredient.id}"></div>
                     </li>`
    ).join('')}
                    ${data.pop().name !== search ? `<li class="ingredient-list-element" id="ingredient-list-id-0">
                        New: ${search}
                        <input type="checkbox" id="ingredient-list-checkbox-id-0" value="0" data-name="${search}" onchange="handleTick(this)">
                        <div id="ingredient-list-details-id-0"></div>
                        </li>` : ""
                    }
    </ul>`
}

var ingredients = []
function handleTick(tick) {
    console.log(tick)
    const elem = document.getElementById(`ingredient-list-details-id-${tick.value}`)
    if (tick.checked) {
        elem.innerHTML = `<input id="ingredient-list-amount-id-${tick.value}" placeholder="Enter amount" class="">
            <input id="ingredient-list-unit-id-${tick.value}" placeholder="Enter unit" class="">
            <input type="button" id="add-ingredient-id-${tick.value}" value="Add ingredient" data-id=${tick.value} data-name="${tick.dataset.name}" onclick="appendIngredient(this)" >`
    } else {
        elem.innerHTML = ""
    }
}

function appendIngredient(input) {
    const id = Number(input.dataset.id)
    const amount = Number(document.getElementById(`ingredient-list-amount-id-${id}`).value)
    const name = input.dataset.name
    const givenUnit = document.getElementById(`ingredient-list-unit-id-${id}`).value
    const unit = givenUnit === "" ? "units" : givenUnit
    ingredients.push({
        id: id,
        name: name,
        amount: amount,
        unit: unit
    })
    document.getElementById(`ingredient-list-details-id-${id}`).innerHTML = ""
    document.getElementById("ingredient-list-search").innerHTML = ""
    document.getElementById("ingredient-search").value = ""
    document.getElementById("new-recipe-ingredients").innerHTML += `<li>${amount} ${unit} ${name}</li>`
}

function addRecipe() {
    const name = document.getElementById("new-recipe-name").value
    const instructions = document.getElementById("new-recipe-instructions").value
    const data = {
        "name": name,
        "text": instructions,
        "ingredients": ingredients
    }
    fetch("../../recipe", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Authorization": "Bearer " + sessionStorage.getItem('token')
        },
        body: JSON.stringify(data),
    }).then(data => {
        console.log(data.body)
        document.getElementById("ingredient-list-search").innerHTML = ""
        document.getElementById("ingredient-search").value = ""
        document.getElementById("new-recipe-name").value = ""
        document.getElementById("new-recipe-ingredients").value = ""
        document.getElementById("new-recipe-instructions").value = ""
        getMyCookies()
    })

}

function checkLoggedIn() {
    fetch("../../check-login", {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            "Authorization": "Bearer " + sessionStorage.getItem('token')
        }
    }).then(res => {
        if (!res.ok) {
            console.log("you do not seem to be logged in")
            alert("Login first")
            document.location.href = '/'
        }
    })

}

function logout() {
    sessionStorage.clear()
    document.location.href = '/'
}