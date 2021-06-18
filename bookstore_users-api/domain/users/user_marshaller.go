package users

import (
	"encoding/json"
)

type PublicUser struct {
	Id          int64  `json:"id"`
	Status      string `json:"status"`
	DateCreated string `json:"date_created"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Status      string `json:"status"`
	DateCreated string `json:"date_created"`
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, u := range users {
		result[index] = u.Marshall(isPublic)
	}
	return result
}

func (u *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:          u.Id,
			DateCreated: u.DateCreated,
			Status:      u.Status,
		}
	}

	userJson, _ := json.Marshal(u)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}
