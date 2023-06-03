package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // Importation du pilote SQLite
)

func main() {
	// Chemin vers la base de données SQLite
	dbPath := "../db/forum.db"

	// Ouverture de la connexion à la base de données
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la base de données:", err)
		return
	}
	defer db.Close()

	// Vérification de la connexion à la base de données
	err = db.Ping()
	if err != nil {
		fmt.Println("Erreur lors de la vérification de la connexion à la base de données:", err)
		return
	}

	fmt.Println("Connexion à la base de données SQLite établie avec succès !")
}
