package repository

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type BankRepository interface {
	GetAllBank(ctx context.Context) (*[]dto.BankResponse, error)
}
