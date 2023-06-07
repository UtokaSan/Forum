package cmd

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/gorilla/sessions"
	"net/http"
)

func authUserSecurity(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		// Gérer le cas où le jeton Bearer n'est pas présent dans l'en-tête
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Extraire le jeton Bearer du header (il sera dans le format "Bearer <token>")
	bearerToken := authHeader[len("Bearer "):]

	fmt.Println(bearerToken)

	// JWT à décoder
	// Clé secrète utilisée pour signer le JWT
	secretKey := []byte("token-user")

	// Parsage du JWT
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		fmt.Println("merde 3")
	}

	// Vérification de la validité du JWT
	if !token.Valid {
		fmt.Println("merde")
		return
	}

	// Accès aux claims (informations) du JWT
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("merde 1")
	}

	// Accès aux données spécifiques du JWT
	userID := claims["user_id"].(string)
	expiration := claims["exp"].(float64)

	// Utilisation des informations extraites du JWT
	fmt.Println("User ID:", userID)
	fmt.Println("Expiration:", expiration)
}
