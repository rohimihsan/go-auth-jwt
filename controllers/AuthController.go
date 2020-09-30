package controllers

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
	"os/user"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"gopkg.in/go-playground/validator.v9"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/rohimihsan/go-auth-jwt/config/db"
	"github.com/rohimihsan/go-auth-jwt/models"
)

var jwtKey = []byte("mEK8VICqacKS0Cy6Ga7vPb2g93SXVZIfsJzrWVQVH64MQRyizWyMGK2E2ugAJ6n")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

//Login function
func Login(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	var res models.ResponseResult

	err := json.NewDecoder(r.Body).Decode(&creds)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db, _ := db.Con()
	var result models.User

	//search for user with email
	email_filter := bson.D{{"email", creds.Email}}

	err = db.Collection("users").FindOne(context.TODO(), email_filter).Decode(&result)

	if result.Email != creds.Email {
		res.Error = err.Error()
		res.Result = "Email Not found"
		res.Data = result
		json.NewEncoder(w).Encode(res)
		return
	}

	match := CheckPasswordHash(creds.Password, result.Password)

	if !match {
		res.Result = "Email and Password does not match"
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(res)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &models.Claims{
		Email: creds.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func Register(w http.ResponseWriter, r *http.Request) {
	a := models.User{
		Email:     r.FormValue("email"),
		Password:  r.FormValue("password"),
		Firstname: r.FormValue("firstname"),
		Lastname:  r.FormValue("lastname"),
	}

	var res models.ResponseResult

	v := validator.New()
	err := v.Struct(a)

	if err != nil {
		res.Error = err.Error()
		res.Data = err.(validator.ValidationErrors)

		json.NewEncoder(w).Encode(res)
		return
	}

	db, _ := db.Con()
	var result models.User

	//check for email duplicate
	email_filter := bson.D{{"email", a.Email}}

	err = db.Collection("users").FindOne(context.TODO(), email_filter).Decode(&result)

	if result.Email == a.Email {
		res.Error = err.Error()
		res.Result = "Email already registered"
		res.Data = result
		json.NewEncoder(w).Encode(res)
		return
	}

	//generate username
	uname := UnameGenerator(a.Firstname + "." + a.Lastname)

	hash, err := HashPassword(a.Password)

	if err != nil {
		res.Error = err.Error()
		res.Result = "Error when trying to hash password"

		json.NewEncoder(w).Encode(res)
		return
	}

	var user_data = bson.D{
		{"firstname", a.Firstname},
		{"lastname", a.Lastname},
		{"username", uname},
		{"email", a.Email},
		{"password", hash},
		{"created_at", time.Now()},
	}

	insertResult, err := db.Collection("users").InsertOne(context.TODO(), user_data)

	if err != nil {
		res.Error = err.Error()
		res.Result = "Error when trying to store data"

		json.NewEncoder(w).Encode(res)
		return
	}

	res.Result = "Success creating account"
	res.Data = insertResult

	json.NewEncoder(w).Encode(res)
	return

}

//refresh jwt token
func RefreshToken(w http.ResponseWriter, r *http.Request) {
	//get cookie
	c, err := r.Cookie("Token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
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
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "Token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

//utility
func UnameGenerator(username string) string {
	random := strconv.Itoa(rand.Intn(9999))

	name := []string{username, random}

	new := strings.Join(name, "")

	//get db
	db, _ := db.Con()

	//check if username exist
	uname_filter := bson.D{{"username", username}}

	var result user.User

	db.Collection("users").FindOne(context.TODO(), uname_filter).Decode(&result)

	for result.Username == new {
		db.Collection("users").FindOne(context.TODO(), uname_filter).Decode(&result)

		random := strconv.Itoa(rand.Intn(9999))

		name := []string{username, random}

		new = strings.Join(name, ".")
	}

	return new
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RandString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}
