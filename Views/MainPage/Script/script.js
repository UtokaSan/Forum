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
