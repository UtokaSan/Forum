function submitForm (event) {
    event.preventDefault();
    // Faire pour tous les users
    const selectRole = document.querySelector("#role-user").value;
    const selectBanUser = document.querySelector("#deban-user").value;
    let formData = {}
    if (selectRole !== "") {
        formData["user-role"] = selectRole
    }
    if (selectBanUser !== "") {
        formData["deban-user"] = selectBanUser
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
        const selectDeban = document.querySelector("#deban-user")
        const selectRoleAdmin = document.querySelector("#role-admin-user");
        const selectRoleModo = document.querySelector("#role-modo-user");
        for (let i = 0; i < data.ban.length; i++) {
            const option = document.createElement("option")
            option.text = data.ban[i].username;
            selectDeban.appendChild(option)
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