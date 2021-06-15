package services

import (
	"github.com/mrmtsu/go-book/bookstore_users-api/domain/users"
	"github.com/mrmtsu/go-book/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func CreateUser(u users.User) (*users.User, *errors.RestErr) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.Save(); err != nil {
		return nil, err
	}

	return &u, nil
}

func UpdateUser(isPartial bool, u users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(u.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if u.FirstName != "" {
			current.FirstName = u.FirstName
		}

		if u.LastName != "" {
			current.LastName = u.LastName
		}

		if u.Email != "" {
			current.Email = u.Email
		}
	} else {
		current.FirstName = u.FirstName
		current.LastName = u.LastName
		current.Email = u.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}
