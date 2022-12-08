package impl

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (s *customerService) UpdateCustomer(ctx context.Context, userID *int, newData *dto.CustomerUpdateRequest) error {

	err := s.repo.UpdateCustomer(ctx, userID, newData)
	if err != nil {
		if err == customerrors.ErrRecordNotFound {
			return err
		}
		s.log.Warningln("[UpdateCustomer] Failed on update to the database:", err.Error())
		return err
	}

	return nil
}
