package usecase

import (
	"Todo/pkg/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	Repo repo.UserRepoInterfaces
}

func NewUserUsecase(Repo repo.UserRepoInterfaces) use.UserUseCaseInterface {
	return &UserUsecase{Repo}
}

func (uu *UserUsecase) ExecuteSignup(request models.User) (*models.User, error) {
	Email, err := uu.UserRepo.FetchEMail(request.Email)
	if err != nil {
		return nil, err
	}
	if Email.email != "" {
		return nil, errors.New("Email already exists")
	}
	Phone, err := uu.UserRepo.FetchPhoneNumber(request.Phone)
	if err != nil {
		return nil, err
	}
	if Phone.phone != "" {
		return nil, errors.New("Phone number already exits")
	}
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	NewUser := &models.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Phone:     request.Phone,
		Email:     request.Email,
		Password:  string(HashedPassword),
	}
	err = uu.UserRepo.CreateUser(NewUser)
	if err != nil {
		return nil, err
	}
	return NewUser, nil
}
