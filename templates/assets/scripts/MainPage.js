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

document.addEventListener("DOMContentLoaded", function (event) {
    var defaultCategoryButton = document.querySelector(".onglet[data-category='contenu1']");
    defaultCategoryButton.classList.add("active");
    var defaultCategoryContent = document.querySelector(".contenu-onglet[data-content='contenu1']");
    defaultCategoryContent.style.display = "block";
    var categoryButtons = document.querySelectorAll(".onglet");
    categoryButtons.forEach(function (button) {
        button.addEventListener("click", function () {
            var category = button.getAttribute("data-category");
            var discussionDivs = document.querySelectorAll(".contenu-onglet");
            discussionDivs.forEach(function (div) {
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
            newDiv.className = "post-div topics " + post.categorie.split(" ").join("_") ;
            // Créer un élément d'ancre
            var anchorElement = document.createElement("a");
            anchorElement.textContent = post.title;
            if (anchorElement.textContent.length > 50) {
                anchorElement.textContent = anchorElement.textContent.substring(0, 47)
                anchorElement.textContent += "..."
            }
            newDiv.appendChild(anchorElement);
            let categoryDiv
            console.log(newDiv.classList.contains('Drogue'))
            if (newDiv.classList.contains('Chat_Général')) {
                document.getElementById("contenu1").appendChild(newDiv);
            } else if (newDiv.classList.contains('Drogue')) {
                document.getElementById("contenu2").appendChild(newDiv);
            } else if (newDiv.classList.contains('Sex_Cam')) {
                document.getElementById("contenu3").appendChild(newDiv);
            } else if (newDiv.classList.contains('Red_Room')) {
                document.getElementById("contenu4").appendChild(newDiv);
            } else {
                document.getElementById("contenu1").appendChild(newDiv);
            }
            newDiv.addEventListener('click', function() {
                let urlParameter = new URLSearchParams();
                urlParameter.append('id', post.id);
                let urlNewPage = `http://localhost:8080/post?${urlParameter.toString()}`;
                window.location.href = urlNewPage;
            });
        });
    })
    .catch(error => {
        console.error("Error update", error);
    });


document.getElementById("createPost").addEventListener("submit", function(event) {
    event.preventDefault();

    var action = document.getElementById("menu-deroulant-select").value;
    var message = document.getElementById("message").value;
    var messageContent = document.getElementById("messageContent").value;
    var imageUpload = document.getElementById("imageUpload").files[0];

    var formData = new FormData();
    formData.append("action", action);
    formData.append("message", message);
    formData.append("messageContent", messageContent);
    formData.append("imageUpload", imageUpload);

    fetch("/api/create-post", {
        method: "POST",
        body: formData
    })
        .then(function(response) {
            if (response.ok) {
                alert("Post created")
                location.reload();
            } else {
                console.error(response.status);
            }
        })
        .catch(function(error) {
            console.error(error);
        });
});