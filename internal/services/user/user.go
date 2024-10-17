package user

import (
	"github.com/banking-system/internal/cache"
	"github.com/banking-system/internal/repository"
)

type UserService struct {
	repo  repository.IDatabase
	cache cache.ICache
}

func NewUserService(repoi repository.IDatabase, ci cache.ICache) *UserService {
	return &UserService{
		repo:  repoi,
		cache: ci,
	}
}

func (u *UserService) Login() string {
	return u.repo.Login()
}

func (u *UserService) RetrieveData() string {
	return u.cache.RetrieveData()
}
