package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gozaddy/AuthProject/database"
	"github.com/gozaddy/AuthProject/models"
	"github.com/gozaddy/AuthProject/utils"
)

var (
	tpl *template.Template
)

func init() {
	tpl = template.Must(template.ParseGlob("views/*.gohtml"))
}

//IndexPage controller
func IndexPage(w http.ResponseWriter, r *http.Request) {
	var u models.User
	var err error
	loggedIn, email := database.AlreadyLoggedIn(w, r)
	if loggedIn {
		u, err = database.GetUser(email)
		fmt.Println(string(email))
		utils.CheckError(w, err)
		tpl.ExecuteTemplate(w, "index.gohtml", u)

	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

}

//LoginPage controller
func LoginPage(w http.ResponseWriter, r *http.Request) {
	loggedIn, _ := database.AlreadyLoggedIn(w, r)
	if loggedIn {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		tpl.ExecuteTemplate(w, "login.gohtml", nil)
	}

}

//SignInPage controller
func SignInPage(w http.ResponseWriter, r *http.Request) {
	loggedIn, _ := database.AlreadyLoggedIn(w, r)
	if loggedIn {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		tpl.ExecuteTemplate(w, "signin.gohtml", nil)
	}

}
