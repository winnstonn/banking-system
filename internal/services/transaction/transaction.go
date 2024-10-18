package transaction

import (
	"github.com/banking-system/internal/cache"
	"github.com/banking-system/internal/repository"
)

type ITransactionService interface {
	Transfer() string
}

type TransactionService struct {
	repo  repository.IDatabase
	cache cache.ICache
}

func NewTransationService(repoi repository.IDatabase, ci cache.ICache) *TransactionService {
	return &TransactionService{
		repo:  repoi,
		cache: ci,
	}
}
