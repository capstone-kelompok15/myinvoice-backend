package service

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

type BankService interface {
	GetAllBank(ctx context.Context) (*[]dto.BankResponse, error)
}
