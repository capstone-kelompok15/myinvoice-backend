package impl

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (s *bankService) GetAllBank(ctx context.Context) (*[]dto.BankResponse, error) {
	bankResponse, err := s.repo.GetAllBank(ctx)

	if err != nil {
		if err != customerrors.ErrRecordNotFound {
			s.log.Warningln("[GetAllBank] Error while get all bank repo", err.Error())
		}
		return nil, customerrors.ErrRecordNotFound
	}

	return bankResponse, nil
}
