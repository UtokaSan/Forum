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

document.addEventListener("DOMContentLoaded", function(event) {
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
        data.forEach(post => {
            var newDiv = document.createElement("div");
            newDiv.id = "post-" + post.id;
            newDiv.className = "post-div topics";

            // Créer un élément d'ancre
            var anchorElement = document.createElement("a");
            anchorElement.href = '/'+post.id;
            anchorElement.textContent = post.title;

            if (anchorElement.textContent.length > 50)
            {
                anchorElement.textContent = anchorElement.textContent.substring(0,47)
                anchorElement.textContent += "..."
            }

            newDiv.appendChild(anchorElement);

            var categoryDiv = document.getElementById("contenu" + (post.id + 1));
            if (categoryDiv) {
                categoryDiv.appendChild(newDiv);
            }

        });
    })
    .catch(error => {
        console.error("Error update", error);
    });
