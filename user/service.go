package user

import (
	"fmt"
	"net/http"
)

type IService interface {
	Save(input InputNewUser) (User, int, error)
	Update(input InputNewUser, id int) (User, int, error)
	GetAll() ([]User, int, error)
	GetByEmail(email string) (User, int, error)
	Delete(id int) (int, error)
}

type service struct {
	repoUser IRespository
}

func NewService(repoUser IRespository) *service {
	return &service{repoUser}
}

func (s *service) Save(input InputNewUser) (User, int, error) {
	userByEmail, err := s.repoUser.FindByEmail(input.Email)
	if err != nil {
		return userByEmail, http.StatusInternalServerError, err
	}

	if userByEmail.ID != 0 {
		return userByEmail, http.StatusBadRequest, fmt.Errorf("user dengan email %v sudah terdaftar", input.Email)
	}

	newUser := User{
		Email:    input.Email,
		Password: input.Password,
		Role:     "user",
	}

	userSaved, err := s.repoUser.Save(newUser)
	if err != nil {
		return userSaved, http.StatusInternalServerError, err
	}

	return userSaved, http.StatusOK, nil
}

func (s *service) Update(input InputNewUser, id int) (User, int, error) {
	userByID, err := s.repoUser.FindByID(id)
	if err != nil {
		return userByID, http.StatusInternalServerError, err
	}

	if userByID.ID == 0 {
		return userByID, http.StatusBadRequest, fmt.Errorf("user dengan id %v tidak terdaftar", id)
	}

	if userByID.Email == input.Email {
		return userByID, http.StatusBadRequest, fmt.Errorf("user dengan email %v sudah terdaftar", input.Email)

	}

	userByID.Email = input.Email
	userByID.Password = input.Password

	userUpdated, err := s.repoUser.Update(userByID)
	if err != nil {
		return userUpdated, http.StatusInternalServerError, err
	}

	return userUpdated, http.StatusOK, nil
}

func (s *service) GetAll() ([]User, int, error) {
	users, err := s.repoUser.FindAll()
	if err != nil {
		return users, http.StatusInternalServerError, err
	}

	return users, http.StatusOK, nil
}

func (s *service) GetByEmail(email string) (User, int, error) {
	userByEmail, err := s.repoUser.FindByEmail(email)
	if err != nil {
		return userByEmail, http.StatusInternalServerError, err
	}

	if userByEmail.ID == 0 {
		return userByEmail, http.StatusBadRequest, fmt.Errorf("user dengan email %v tidak terdaftar", email)
	}
	return userByEmail, http.StatusOK, nil
}

func (s *service) Delete(id int) (int, error) {
	userByID, err := s.repoUser.FindByID(id)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	if userByID.ID == 0 {
		return http.StatusBadRequest, fmt.Errorf("user dengan id %v tidak terdaftar", id)
	}

	if err := s.repoUser.Delete(id); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
