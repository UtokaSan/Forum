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
