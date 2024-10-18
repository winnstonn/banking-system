package repository

import (
	"context"

	"github.com/banking-system/internal/models"
)

func (r *Repo) Login() string {
	return "DB Data"
}

func (r *Repo) Transfer(ctx context.Context, transferRecord *models.TransferRecord) (string, error) {
	tx, errCreateTx := r.dbConn.BeginTx(ctx, nil)
	if errCreateTx != nil {
		return "", errCreateTx
	}
	defer tx.Rollback()

	tx.Commit()
	return "Successful Insert Record", nil
}
