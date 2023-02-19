package controllers

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/husnain3184/manup-master/entities"
	"github.com/husnain3184/manup-master/models"
)

var speakersModel = models.NewSpeaker()

func Index(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{
		"data": template.HTML(GetData()),
	}

	temp, _ := template.ParseFiles("views/frontend/index.html", "views/frontend/headerfooter.html")
	temp.Execute(w, data)
}

func GetData() string {
	// fmt.Println("test")
	buffer := &bytes.Buffer{}
	temp, _ := template.New("speakers.html").Funcs(template.FuncMap{
		"increment": func(a, b int) int {
			return a + b
		},
	}).ParseFiles("views/frontend/speakers.html")
	var speakers []entities.Speakers
	err := speakersModel.FindAll(&speakers)
	if err != nil {
		panic(err)
	}
	data := map[string]interface{}{
		"speakers": speakers,
	}
	temp.ExecuteTemplate(buffer, "speakers.html", data)
	return buffer.String()
}

// func About(w http.ResponseWriter, r *http.Request) {
// 	temp, _ := template.ParseFiles("views/frontend/About.html", "views/frontend/headerfooter.html")
// 	temp.Execute(w, nil)
// }
// func Contact(w http.ResponseWriter, r *http.Request) {
// 	temp, _ := template.ParseFiles("views/frontend/contact-us.html", "views/frontend/headerfooter.html")
// 	temp.Execute(w, nil)
// }
// func Speaker(w http.ResponseWriter, r *http.Request) {
// 	temp, _ := template.ParseFiles("views/frontend/Speaker.html", "views/frontend/headerfooter.html")
// 	temp.Execute(w, nil)
// }
// func Schedule(w http.ResponseWriter, r *http.Request) {
// 	temp, _ := template.ParseFiles("views/frontend/Schedule.html", "views/frontend/headerfooter.html")
// 	temp.Execute(w, nil)
// }
// func Blog(w http.ResponseWriter, r *http.Request) {
// 	temp, _ := template.ParseFiles("views/frontend/Blog.html", "views/frontend/headerfooter.html")
// 	temp.Execute(w, nil)
// }
// func BlogDetail(w http.ResponseWriter, r *http.Request) {
// 	temp, _ := template.ParseFiles("views/frontend/blog-details.html", "views/frontend/headerfooter.html")
// 	temp.Execute(w, nil)
// }
