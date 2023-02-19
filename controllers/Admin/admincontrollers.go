package controllers

import (
	"errors"
	"log"
	"net/http"
	"text/template"

	"github.com/husnain3184/manup-master/config"
	"github.com/husnain3184/manup-master/entities"

	"github.com/husnain3184/manup-master/models"
	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Username string
	Password string
	Email    string
}

var adminModel = models.NewAdminModel()

func AdminLogin(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		temp, _ := template.ParseFiles("views/backend/adminlogin.html")
		temp.Execute(w, nil)

	} else if r.Method == http.MethodPost {

		r.ParseForm()

		UserInput := &UserInput{

			Username: r.Form.Get("username"),
			Password: r.Form.Get("password"),
		}

		var user entities.Admin
		// fmt.Println(user)
		adminModel.Where(&user, "username", UserInput.Username)

		var message error
		if UserInput.Username == "" {

			message = errors.New("username and password required")
		} else {

			errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserInput.Password))
			log.Println(errPassword)

			if errPassword != nil {

				message = errors.New("username and password not matched")
			}

		}

		if message != nil {

			data := map[string]interface{}{

				"error": message,
			}
			temp, _ := template.ParseFiles("views/backend/adminlogin.html")
			temp.Execute(w, data)
		} else {

			session, _ := config.Store.Get(r, config.SESSION_ID)

			session.Values["loggedIn"] = true
			session.Values["email"] = user.Email
			session.Values["username"] = user.Name
			session.Values["name"] = user.Name

			session.Save(r, w)

			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)

		}

	}

}

func Dashboard(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)
	if len(session.Values) == 0 {

		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	} else {

		if session.Values["loggedIn"] != true {

			http.Redirect(w, r, "/admin", http.StatusSeeOther)
		} else {

			data := map[string]interface{}{

				"name": session.Values["name"],
			}

			temp, _ := template.ParseFiles("views/backend/dashboard.html", "views/backend/dashboardlayout.html")
			temp.Execute(w, data)

		}

	}

}

func Logout(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)

	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/admin", http.StatusSeeOther)
}

func Register(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		temp, _ := template.ParseFiles("views/backend/adminregister.html")
		temp.Execute(w, nil)

	} else if r.Method == http.MethodPost {

		r.ParseForm()

		user := entities.Admin{

			Name:      r.Form.Get("name"),
			Email:     r.Form.Get("email"),
			Username:  r.Form.Get("username"),
			Password:  r.Form.Get("password"),
			CPassword: r.Form.Get("cpassword"),
		}

		errorMessages := make(map[string]interface{})

		if user.Name == "" {

			errorMessages["Name"] = "Name Required"
		}

		if user.Email == "" {

			errorMessages["Email"] = "Email Required"
		}

		if user.Username == "" {

			errorMessages["Username"] = "Username Required"
		}

		if user.Password == "" {

			errorMessages["Password"] = "Password Required"
		}

		if user.CPassword == "" {

			errorMessages["CPassword"] = "Confirm Password Required"
		} else {

			if user.Password != user.CPassword {

				errorMessages["CPassword"] = "Confirm password not match with password"
			}
		}

		if len(errorMessages) > 0 {

			data := map[string]interface{}{

				"validation": errorMessages,
			}

			temp, _ := template.ParseFiles("views/adminregister.html")
			temp.Execute(w, data)

		} else {

			hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			user.Password = string(hashPassword)
			// fmt.Println(user.Password)

			_, err := adminModel.Create(user)
			var message string
			if err != nil {

				message = "Register Not successfully" + message
			} else {

				message = "Register Success"
			}

			data := map[string]interface{}{

				"success": message,
			}

			temp, _ := template.ParseFiles("views/backend/adminregister.html")
			temp.Execute(w, data)

		}
	}
}
