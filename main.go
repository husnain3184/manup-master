package main

import (
	"net/http"

	admincontroller "github.com/husnain3184/manup-master/controllers/Admin"
	"github.com/husnain3184/manup-master/controllers/Admin/managecoursecontroller"
	"github.com/husnain3184/manup-master/controllers/Admin/managespeakerscontroller"
	homecontroller "github.com/husnain3184/manup-master/controllers/Frontend/home"
)

func main() {
	http.HandleFunc("/", homecontroller.Index)
	// http.HandleFunc("/about-us", homecontroller.About)
	// http.HandleFunc("/contact-us", homecontroller.Contact)
	// http.HandleFunc("/speaker", homecontroller.Speaker)
	// http.HandleFunc("/schedule", homecontroller.Schedule)
	// http.HandleFunc("/blog", homecontroller.Blog)
	// http.HandleFunc("/blog-details", homecontroller.BlogDetail)
	http.HandleFunc("/admin", admincontroller.AdminLogin)
	http.HandleFunc("/register", admincontroller.Register)
	http.HandleFunc("/dashboard", admincontroller.Dashboard)
	http.HandleFunc("/manage-course", managecoursecontroller.Index)
	http.HandleFunc("/course/store", managecoursecontroller.Store)
	http.HandleFunc("/course/get_form", managecoursecontroller.GetForm)
	http.HandleFunc("/course/delete", managecoursecontroller.Delete)
	http.HandleFunc("/manage-speakers", managespeakerscontroller.Index)
	http.HandleFunc("/manage-speakers/store", managespeakerscontroller.Store)
	http.HandleFunc("/manage-speakers/get_form", managespeakerscontroller.GetForm)
	http.HandleFunc("/manage-speakers/delete", managespeakerscontroller.Delete)
	http.HandleFunc("/logout", admincontroller.Logout)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8000", nil)
}
