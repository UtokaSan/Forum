function afficherContenu(evt, contenuId) {
    // Cacher tous les contenus des onglets
    var contenus = document.getElementsByClassName("contenu-onglet");
    for (var i = 0; i < contenus.length; i++) {
        contenus[i].style.display = "none";
    }

    // Supprimer la classe "active" de tous les onglets
    var onglets = document.getElementsByClassName("onglet");
    for (var i = 0; i < onglets.length; i++) {
        onglets[i].className = onglets[i].className.replace(" active", "");
    }

    // Afficher le contenu de l'onglet sélectionné
    document.getElementById(contenuId).style.display = "block";

    // Ajouter la classe "active" à l'onglet sélectionné
    evt.currentTarget.className += " active";
}

document.addEventListener("DOMContentLoaded", function(event) {
    // Activer la catégorie "Chat Général" et afficher son contenu
    var defaultCategoryButton = document.querySelector(".onglet[data-category='contenu1']");
    defaultCategoryButton.classList.add("active");
    var defaultCategoryContent = document.querySelector(".contenu-onglet[data-content='contenu1']");
    defaultCategoryContent.style.display = "block";

    var categoryButtons = document.querySelectorAll(".onglet");

    categoryButtons.forEach(function(button) {
        button.addEventListener("click", function() {
            var category = button.getAttribute("data-category");

            var discussionDivs = document.querySelectorAll(".contenu-onglet");

            discussionDivs.forEach(function(div) {
                if (div.getAttribute("data-content") === category) {
                    div.style.display = "block";
                } else {
                    div.style.display = "none";
                }
            });
        });
    });
});


fetch("/api/display-post", {
    method: "POST",
    headers: {
        "Content-Type": "application/json"
    },
})
    .then(response => response.json())
        .then(data => {
            console.log(data)
        })
    .catch(error => {
        console.error("Error update", error);
    });

fetch("/api/display-post", {
    method: "POST",
    headers: {
        "Content-Type": "application/json"
    },
})
    .then(response => response.json())
    .then(data => {

        data.forEach(post => {

            var newDiv = document.createElement("div");
            newDiv.id = "post-" + post.id; // Utiliser l'ID du post comme identifiant de la div
            newDiv.className = "post-div";

            var titleElement = document.createElement("h2");
            titleElement.textContent = post.title;

            newDiv.appendChild(titleElement);

            var categoryDiv = document.getElementById("contenu" + (post.id + 1));
            if (categoryDiv) {
                categoryDiv.appendChild(newDiv);
            }
        });
    })
    .catch(error => {
        console.error("Error update", error);
    });
