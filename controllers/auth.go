package controllers

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/gozaddy/AuthProject/models"

	"github.com/gozaddy/AuthProject/utils"

	"github.com/gozaddy/AuthProject/database"
)

//Register user
func Register(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	p := r.FormValue("password")
	//rememberUser := r.FormValue("remember")
	//fmt.Println(rememberUser)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	utils.CheckError(w, err)

	u := models.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	err = database.StoreUser(u)
	utils.CheckError(w, err)

	sessionID, err := uuid.NewV4()
	utils.CheckError(w, err)

	http.SetCookie(w, &http.Cookie{
		Name:     "session-id",
		Value:    sessionID.String(),
		HttpOnly: true,
	})

	_, err = database.Cache.Do("SETEX", sessionID.String(), "300", u.Email)

	utils.CheckError(w, err)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

//Login user
func Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	p := r.FormValue("password")

	u, err := database.GetUser(email)
	if err != nil {
		http.Error(w, "User does not exist", http.StatusNotFound)
		return
	}

	err = bcrypt.CompareHashAndPassword(u.Password, []byte(p))
	if err != nil {
		http.Error(w, "Incorrect password", http.StatusUnauthorized)
	}

	sessionID, err := uuid.NewV4()
	utils.CheckError(w, err)

	http.SetCookie(w, &http.Cookie{
		Name:     "session-id",
		Value:    sessionID.String(),
		HttpOnly: true,
	})

	_, err = database.Cache.Do("SETEX", sessionID.String(), "300", u.Email)
	utils.CheckError(w, err)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

//LogOut user
func LogOut(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session-id")
	utils.CheckError(w, err)
	_, err = database.Cache.Do("DEL", c.Value)
	utils.CheckError(w, err)
	c = &http.Cookie{
		Name:   "",
		Value:  "",
		MaxAge: -1,
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)

}
