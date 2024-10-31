package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/banking-system/internal/models"
)

func TestRepo_Transfer(t *testing.T) {
	type fields struct {
		dbConn *sql.DB
	}
	type args struct {
		ctx            context.Context
		transferRecord *models.TransferRecord
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{
				dbConn: tt.fields.dbConn,
			}
			got, err := r.Transfer(tt.args.ctx, tt.args.transferRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repo.Transfer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Repo.Transfer() = %v, want %v", got, tt.want)
			}
		})
	}
}
