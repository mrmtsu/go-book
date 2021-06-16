package users

import (
	"fmt"

	"github.com/mrmtsu/go-book/bookstore_users-api/datasources/mysql/users_db"
	"github.com/mrmtsu/go-book/bookstore_users-api/utils/errors"
	"github.com/mrmtsu/go-book/bookstore_users-api/utils/mysql_utils"
)

const (
	errorNoRows           = "no rows in result set"
	queryInsertUser       = "INSERT INTO users_db.users(first_name, last_name, email, status, password, date_created) VALUES(?, ?, ?, ?, ?, ?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, status, date_created FROM users_db.users WHERE id=?;"
	queryUpdateUser       = "UPDATE users_db.users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM userd_db.users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, firt_name, last_name, email, status, date_created FROM users_db.users WHERE status=?;"
)

func (u *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(u.Id)

	if getErr := result.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.Status, &u.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}
	return nil
}

func (u *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(u.FirstName, u.LastName, u.Email, u.Status, u.Password, u.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}

	userId, saveErr := insertResult.LastInsertId()
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}
	u.Id = userId
	return nil
}

func (u *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (u *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(u.Id); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil
}

func (u *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated, &u.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no uses matching status %s", status))
	}
	return results, nil
}
