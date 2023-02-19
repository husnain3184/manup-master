package models

import (
	"database/sql"

	"github.com/husnain3184/manup-master/config"
	"github.com/husnain3184/manup-master/entities"
)

type CourseModel struct {
	db *sql.DB
}

func New() *CourseModel {

	db, err := config.DBConnection()
	if err != nil {

		panic(err)
	}
	return &CourseModel{db: db}
}

func (m *CourseModel) FindAll(course *[]entities.Course) error {

	rows, err := m.db.Query("select * from course")
	if err != nil {

		return err
	}

	defer rows.Close()

	for rows.Next() {

		var data entities.Course
		rows.Scan(&data.Id, &data.CourseName, &data.Lesson, &data.Week, &data.Price, &data.Description)
		*course = append(*course, data)
	}
	return nil
}

func (m *CourseModel) Find(id int64, course *entities.Course) error {

	return m.db.QueryRow("select * from course where id = ?", id).Scan(

		&course.Id,
		&course.CourseName,
		&course.Lesson,
		&course.Week,
		&course.Price,
		&course.Description)
}

func (m *CourseModel) Create(course *entities.Course) error {

	result, err := m.db.Exec("insert into course (coursename,lesson,week,price,description)values(?,?,?,?,?)", course.CourseName, course.Lesson, course.Week, course.Price, course.Description)
	if err != nil {

		return err
	}

	lastInsertId, _ := result.LastInsertId()
	course.Id = lastInsertId
	return nil

}

func (m *CourseModel) Update(course entities.Course) error {

	_, err := m.db.Exec("update course set coursename = ?, lesson = ?, week = ?, price = ?, description = ? where id = ?", course.CourseName, course.Lesson, course.Week, course.Price, course.Description, course.Id)

	if err != nil {
		return err
	}

	return nil
}

func (m *CourseModel) Delete(id int64) error {

	_, err := m.db.Exec("delete from course where id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
