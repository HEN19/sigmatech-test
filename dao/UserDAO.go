package dao

import (
	"database/sql"
	"fmt"

	"github.com/api-skeleton/model"
)

type userDAO struct {
	AbstractDAO
}

var UserDAO = userDAO{}.New()

func (input userDAO) New() (output userDAO) {
	output.TableName = "users"
	output.FileName = "UserDAO.go"
	return
}

func (input userDAO) InsertUser(db *sql.DB, inputStruct model.UserModel) (err error) {
	var (
		query string
	)

	query = fmt.Sprintf(
		`
			INSERT INTO %s
				(username, password, first_name, last_name, gender, phone, email, address)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
		`, input.TableName,
	)

	params := []interface{}{
		inputStruct.Username.String, inputStruct.Password.String, inputStruct.FirstName.String,
		inputStruct.LastName.String, inputStruct.Gender.String, inputStruct.Telephone.String,
		inputStruct.Email.String, inputStruct.Address.String,
	}

	_, err = db.Exec(query, params...)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}

func (input userDAO) LoginCheck(db *sql.DB, user model.UserModel) (result model.UserModel, err error) {
	query := "SELECT id, username, first_name, last_name " +
		" FROM " + input.TableName +
		" WHERE username = $1 AND password = $2 "

	param := []interface{}{user.Username.String, user.Password.String}

	results := db.QueryRow(query, param...)
	dbError := results.Scan(&result.ID, &result.Username, &result.FirstName, &result.LastName)

	if dbError != nil && dbError.Error() != "sql: no rows in result set" {
		err = dbError
		return
	}

	return
}

func (input userDAO) GetUserProfile(db *sql.DB, id int64) (model.UserModel, error) {
	query := "SELECT first_name, last_name, gender, phone, email, address, created_at, updated_at " +
		"FROM " + input.TableName + " WHERE id = $1"

	row := db.QueryRow(query, id)
	var user model.UserModel

	err := row.Scan(&user.FirstName, &user.LastName,
		&user.Gender, &user.Telephone, &user.Email, &user.Address,
		&user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}
		return user, err
	}
	return user, nil
}
