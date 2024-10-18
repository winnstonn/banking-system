package controller

import (
	"github.com/banking-system/internal/cache"
	"github.com/banking-system/internal/repository"
	"github.com/banking-system/internal/services/transaction"
	"github.com/banking-system/internal/services/user"
)

type Controller struct {
	userService        user.IUserService
	transactionService transaction.ITransactionService
}

func Init(repository repository.IDatabase, cache cache.ICache) *Controller {
	us := user.NewUserService(repository, cache)
	tr := transaction.NewTransationService(repository, cache)

	return &Controller{
		userService:        us,
		transactionService: tr,
	}
}
