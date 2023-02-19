package models

import (
	"database/sql"

	"github.com/husnain3184/manup-master/config"
	"github.com/husnain3184/manup-master/entities"
)

type AdminModel struct {
	db *sql.DB
}

func NewAdminModel() *AdminModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}
	return &AdminModel{
		db: conn,
	}
}
func (u AdminModel) Where(admin *entities.Admin, fieldName, fieldValue string) error {
	row, err := u.db.Query("select * from admin where "+fieldName+"= ? limit 1", fieldValue)
	// log.Println(row)

	if err != nil {
		return err
	}
	defer row.Close()
	for row.Next() {
		row.Scan(&admin.Id, &admin.Name, &admin.Email, &admin.Username, &admin.Password)
	}
	return nil
}

//	func (u UserModel) Where(user *entities.User, fieldName, fieldValue string) error {
//		row, err := u.db.Query("select * from users where "+fieldName+"= ? limit 1", fieldValue)
//		if err != nil {
//			return err
//		}
//		defer row.Close()
//		for row.Next() {
//			row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
//		}
//		return nil
//	}
func (u AdminModel) Create(admin entities.Admin) (int64, error) {

	result, err := u.db.Exec("insert into admin (name,email,username,password) values (?,?,?,?)", admin.Name, admin.Email, admin.Username, admin.Password)
	if err != nil {
		return 0, err
	}
	lastInsertId, _ := result.LastInsertId()
	return lastInsertId, nil
}
