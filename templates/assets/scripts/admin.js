function afficherContenu(evt, contenuId) {
    var contenus = document.getElementsByClassName("contenu-onglet");
    for (var i = 0; i < contenus.length; i++) {
        contenus[i].style.display = "none";
    }
    var onglets = document.getElementsByClassName("onglet");
    for (var i = 0; i < onglets.length; i++) {
        onglets[i].className = onglets[i].className.replace(" active", "");
    }
    document.getElementById(contenuId).style.display = "block";
    evt.currentTarget.className += " active";
}

function submitForm (event) {
    event.preventDefault();
    const selectUnBanUser = document.querySelector("#unban-user").value;
    const selectBanUser = document.querySelector("#ban-user").value;
    const selectRoleAdmin = document.querySelector("#role-admin-user").value;
    const selectRoleModo = document.querySelector("#role-modo-user").value;
    const selectDeletePost = document.querySelector("#delete-post").value;
    const selectHiddenPost = document.querySelector("#hidden-post").value;

    let formData = {}
    if (selectUnBanUser !== "") {
        formData["key"] = "unban-user"
        formData["unban-user"] = selectUnBanUser
    }
    if (selectBanUser !== "") {
        formData["key"] = "ban-user"
        formData["ban-user"] = selectBanUser
    }
    if (selectRoleModo !== "") {
        formData["key"] = "role-modo-user"
        formData["role-modo-user"] = selectRoleModo
    }
    if (selectRoleAdmin !== "") {
        formData["key"] = "role-admin-user"
        formData["role-admin-user"] = selectRoleAdmin
    }
    if (selectDeletePost !== "") {
        formData["key"] = "delete-post"
        formData["delete-post"] = selectDeletePost
    }
    if (selectHiddenPost !== "") {
        formData["key"] = "hidden-post"
        formData["hidden-post"] = selectHiddenPost
    }
    fetch("/api/adminpanel", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(formData)
    })
        .then(response => {
            console.log("update successfully");
        })
        .catch(error => {
            console.error("Error update", error);
        });
}
// Faire un fetch avec ma fonction adminPanel
fetch("/api/catch-info-admin", {
    method: "POST",
    headers: {
        "Content-Type": "application/json"
    },
    body: JSON.stringify({ /* Vos données à envoyer */ })
})
    .then(response => response.json())
    .then(data => {
        console.log(data)
        const selectDeban = document.querySelector("#unban-user")
        const selectBanUser = document.querySelector("#ban-user");
        const selectRoleAdmin = document.querySelector("#role-admin-user");
        const selectRoleModo = document.querySelector("#role-modo-user");
        const selectDeletePost = document.querySelector("#delete-post");
        const selectHiddenPost = document.querySelector("#hidden-post");
        for (let i = 0; i < data.account.length; i++) {
            const option = document.createElement("option")
            if (data.account[i].ban === 1) {
                option.text = data.account[i].username;
                selectDeban.appendChild(option)
            }
        }
        for (let i = 0; i < data.account.length; i++) {
            const option = document.createElement("option")
            if (data.account[i].ban === 0) {
                option.text = data.account[i].username;
                selectBanUser.appendChild(option)
            }
        }
        for (let i = 0; i < data.postUnHidden.length; i++) {
            const option = document.createElement("option");
            option.text = data.postUnHidden[i].title;
            selectDeletePost.appendChild(option);
        }
        for (let i = 0; i < data.postHidden.length; i++) {
            const option = document.createElement("option");
            option.text = data.postHidden[i].title;
            selectHiddenPost.appendChild(option);
        }
        for (let i = 0; i < data.account.length; i++) {
            const option = document.createElement("option");
            option.text = data.account[i].username;
            const clonedOption = option.cloneNode(true);
            selectRoleAdmin.appendChild(clonedOption);
            const clonedOption2 = option.cloneNode(true);
            selectRoleModo.appendChild(clonedOption2);
        }
    })
    .catch(error => {
        console.error("Erreur lors de la mise à jour :", error);
    });