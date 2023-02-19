package managecoursecontroller

import (
	"bytes"
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/husnain3184/manup-master/entities"
	"github.com/husnain3184/manup-master/models"
)

var courseModel = models.New()

func Index(w http.ResponseWriter, r *http.Request) {

	data := map[string]interface{}{

		"data": template.HTML(GetData()),
	}

	temp, _ := template.ParseFiles("views/backend/course.html", "views/backend/dashboardlayout.html")
	temp.Execute(w, data)
}

func GetData() string {

	buffer := &bytes.Buffer{}
	temp, _ := template.New("data.html").Funcs(template.FuncMap{
		"increment": func(a, b int) int {

			return a + b
		},
	}).ParseFiles("views/backend/data.html")

	var course []entities.Course
	err := courseModel.FindAll(&course)
	if err != nil {

		panic(err)
	}

	data := map[string]interface{}{

		"course": course,
	}

	temp.ExecuteTemplate(buffer, "data.html", data)
	return buffer.String()
}

func GetForm(w http.ResponseWriter, r *http.Request) {

	queryString := r.URL.Query()
	id, err := strconv.ParseInt(queryString.Get("id"), 10, 64)

	var data map[string]interface{}

	if err != nil {

		data = map[string]interface{}{
			"title": "Add Student Data",
		}
	} else {

		var course entities.Course
		err := courseModel.Find(id, &course)
		if err != nil {

			panic(err)
		}
		data = map[string]interface{}{
			"title":  "Edit Student Data",
			"course": course,
		}
	}

	temp, _ := template.ParseFiles("views/backend/form.html")
	temp.Execute(w, data)

}

func Store(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		r.ParseForm()
		var course entities.Course
		course.CourseName = r.Form.Get("coursename")
		course.Lesson = r.Form.Get("lesson")
		course.Week = r.Form.Get("week")
		course.Price = r.Form.Get("price")
		course.Description = r.Form.Get("description")

		id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)

		var data map[string]interface{}

		if err != nil {

			err := courseModel.Create(&course)
			if err != nil {

				RepsonseError(w, http.StatusInternalServerError, err.Error())
				return
			}

			data = map[string]interface{}{

				"message": "Data Successfully Changed",
				"data":    template.HTML(GetData()),
			}
		} else {

			course.Id = id
			err := courseModel.Update(course)
			if err != nil {

				RepsonseError(w, http.StatusInternalServerError, err.Error())
				return
			}
			data = map[string]interface{}{

				"message": "Data Successfully Updated",
				"data":    template.HTML(GetData()),
			}
		}

		ResponseJson(w, http.StatusOK, data)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	id, err := strconv.ParseInt(r.Form.Get("id"), 10, 64)

	if err != nil {

		panic(err)
	}
	err = courseModel.Delete(id)
	if err != nil {

		panic(err)
	}

	data := map[string]interface{}{

		"message": "Student Delete Successfully",
		"data":    template.HTML(GetData()),
	}
	ResponseJson(w, http.StatusOK, data)
}

func RepsonseError(w http.ResponseWriter, code int, message string) {

	ResponseJson(w, code, map[string]string{"error": message})

}

func ResponseJson(w http.ResponseWriter, code int, payload interface{}) {

	// fmt.Println("Payload", payload)
	response, _ := json.Marshal(payload)
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
