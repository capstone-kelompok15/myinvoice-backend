package impl

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (s *customerService) GetSummary(ctx context.Context, customerID int) (*dto.CustomerSummary, error) {
	summary, err := s.repo.GetSummary(ctx, customerID)
	if err != nil {
		if err != customerrors.ErrRecordNotFound {
			s.log.Warningln("[GetSummary] Error on executing the repo:", err.Error())
		}
		return nil, err
	}
	return summary, err
}
