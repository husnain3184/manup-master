package models

import (
	"database/sql"

	"github.com/husnain3184/manup-master/config"
	"github.com/husnain3184/manup-master/entities"
)

type SpeakerModel struct {
	db *sql.DB
}

func NewSpeaker() *SpeakerModel {

	db, err := config.DBConnection()
	if err != nil {

		panic(err)
	}
	return &SpeakerModel{db: db}
}

func (m *SpeakerModel) FindAll(speakers *[]entities.Speakers) error {

	rows, err := m.db.Query("select * from 	speakers")
	if err != nil {

		return err
	}

	defer rows.Close()

	for rows.Next() {

		var data entities.Speakers
		rows.Scan(&data.Id, &data.SpeakerName, &data.Facebook, &data.Instagram, &data.Twiter, &data.Linkdin, &data.Image)
		*speakers = append(*speakers, data)
	}
	return nil
}

func (m *SpeakerModel) Find(id int64, speakers *entities.Speakers) error {

	return m.db.QueryRow("select * from speakers where id = ?", id).Scan(

		&speakers.Id,
		&speakers.SpeakerName,
		&speakers.Facebook,
		&speakers.Instagram,
		&speakers.Twiter,
		&speakers.Linkdin,
		&speakers.Image)
}

func (m *SpeakerModel) Create(speakers *entities.Speakers) error {

	result, err := m.db.Exec("insert into speakers (speakername,facebook,instagram,twiter,linkdin,image)values(?,?,?,?,?,?)", speakers.SpeakerName, speakers.Facebook, speakers.Instagram, speakers.Twiter, speakers.Linkdin, speakers.Image)
	if err != nil {

		return err
	}

	lastInsertId, _ := result.LastInsertId()
	speakers.Id = lastInsertId
	return nil

}

func (m *SpeakerModel) Update(speakers entities.Speakers) error {

	_, err := m.db.Exec("update speakers set speakername = ?, facebook = ?, instagram = ?, twiter = ?, linkdin = ? , image = ? where id = ?", speakers.SpeakerName, speakers.Facebook, speakers.Instagram, speakers.Twiter, speakers.Linkdin, speakers.Image, speakers.Id)

	if err != nil {
		return err
	}

	return nil
}

func (m *SpeakerModel) Delete(id int64) error {

	_, err := m.db.Exec("delete from speakers where id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
