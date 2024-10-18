package user

import (
	"github.com/banking-system/internal/cache"
	"github.com/banking-system/internal/repository"
)

type IUserService interface {
	Login() string
	RetrieveData() string
}

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
