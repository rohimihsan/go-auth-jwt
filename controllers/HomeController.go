package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/rohimihsan/go-auth-jwt/models"
)

//Home function
func Home(w http.ResponseWriter, r *http.Request) {
	var res models.ResponseResult

	// fmt.Fprint(w, "Welcome")
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)

			res.Result = "Unauthorized"
			json.NewEncoder(w).Encode(res)
			return
		}
		w.WriteHeader(http.StatusBadRequest)

		res.Result = "Bad request"
		json.NewEncoder(w).Encode(res)
		return
	}

	tknStr := c.Value

	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			res.Result = "Unauthorized"
			json.NewEncoder(w).Encode(res)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		res.Result = "Unauthorized"
		json.NewEncoder(w).Encode(res)
		return
	}

	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Email)))
}
