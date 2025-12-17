package service

import (
	"session-18/model"
	"session-18/repository"
)

type UserService interface {
	GetUserByID(id int) (model.User, error)
}

type userService struct {
	Repo repository.Repository
}

func NewUserService(repo repository.Repository) UserService {
	return &userService{Repo: repo}
}

func (s *userService) GetUserByID(id int) (model.User, error) {
	return s.Repo.UserRepo.GetUserByID(id)
}
