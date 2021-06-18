package services

import (
	"github.com/mrmtsu/go-book/bookstore_users-api/domain/users"
	"github.com/mrmtsu/go-book/bookstore_users-api/utils/crypto_utils"
	"github.com/mrmtsu/go-book/bookstore_users-api/utils/date_utils"
	"github.com/mrmtsu/go-book/bookstore_users-api/utils/errors"
)

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *usersService) CreateUser(u users.User) (*users.User, *errors.RestErr) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	u.Status = users.StatusActive
	u.DateCreated = date_utils.GetNowDBFormat()
	u.Password = crypto_utils.GetMd5(u.Password)
	if err := u.Save(); err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *usersService) UpdateUser(isPartial bool, u users.User) (*users.User, *errors.RestErr) {
	current, err := s.GetUser(u.Id)
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

func (s *usersService) DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (s *usersService) SearchUser(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
