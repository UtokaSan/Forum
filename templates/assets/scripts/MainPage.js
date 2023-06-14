document.addEventListener("DOMContentLoaded", function(event) {
    var categoryButtons = document.querySelectorAll(".categorie");

    categoryButtons.forEach(function(button) {
        button.addEventListener("click", function() {
            categoryButtons.forEach(function(btn) {
                btn.classList.remove("active");
            });

            button.classList.add("active");

            var category = button.textContent;

            var discussionDivs = document.querySelectorAll(".discussion");

           discussionDivs.forEach(function(div) {
                if (div.classList.contains(category)) {
                    div.classList.remove("hidden");
                } else {
                    div.classList.add("hidden");
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