package utils

import (
	"github.com/go-chi/jwtauth/v5"
	"net/http"
)

var TokenAuth *jwtauth.JWTAuth

func init() {
	TokenAuth = jwtauth.New("HS256", []byte(JWT_SECRET), nil) // replace with secret key
}

func Generate(w http.ResponseWriter) {
	_, token, err := TokenAuth.Encode(map[string]any{"user_id": 123})
	if err != nil {
		http.Error(w, "token error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
    Name:     "jwt",
    Value:    token,
    Path:     "/",
    HttpOnly: true,
})

	w.Write([]byte(token))
}
