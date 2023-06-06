function submitForm (event) {
    event.preventDefault();
    const selectRole = document.querySelector("#role-user").value;
    const selectBanUser = document.querySelector("#deban-user").value;
    let formData = {}
    if (selectRole !== "") {
        formData["user-role"] = selectRole
    }
    if (selectBanUser !== "") {
        formData["deban-user"] = selectBanUser
    }
    fetch("/api/admin", {
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
        console.log(data);
    })
    .catch(error => {
        console.error("Erreur lors de la mise à jour :", error);
    });