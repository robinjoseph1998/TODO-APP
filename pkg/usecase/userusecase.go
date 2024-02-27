package usecase

import (
	"Todo/pkg/models"
	repo "Todo/pkg/repository/interfaces"
	use "Todo/pkg/usecase/interfaces"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	UserRepo repo.UserRepoInterfaces
}

func NewUserUsecase(Repo repo.UserRepoInterfaces) use.UserUseCaseInterface {
	return &UserUsecase{Repo}
}

func (uu *UserUsecase) ExecuteSignup(request models.User) (*models.User, error) {
	Email, err := uu.UserRepo.FetchEmail(request.Email)
	if err != nil {
		return nil, err
	}
	if Email {
		return nil, errors.New("email already exists")
	}
	Phone, err := uu.UserRepo.FetchPhoneNumber(request.Phone)
	if err != nil {
		return nil, err
	}
	if Phone {
		return nil, errors.New("phone number already exits")
	}
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	NewUser := models.User{
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
	return &NewUser, nil
}
func (uu *UserUsecase) ExecuteLogin(phone, password string) (string, error) {
	Phone, err := uu.UserRepo.FetchPhoneNumber(phone)
	if err != nil {
		return "", err
	}
	if !Phone {
		return "", errors.New("user with this phone number not exsits")
	}
	user, err := uu.UserRepo.FetchUser(phone)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("user not exists")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}
	return user.ID.Hex(), nil

}
