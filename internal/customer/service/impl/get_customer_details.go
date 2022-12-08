package impl

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
)

func (s *customerService) GetCustomerDetails(ctx context.Context, req *dto.CustomerContext) (*dto.CustomerDetails, error) {
	customerDetails, err := s.repo.GetCustomerDetail(ctx, req.ID)
	if err != nil {
		if err == customerrors.ErrRecordNotFound {
			return nil, err
		}
		s.log.Warningln("[GetCustomerDetails] Error while getting customer details:", err.Error())
		return nil, err
	}

	if customerDetails.DisplayProfilePictureURL == nil {
		customerDetails.DisplayProfilePictureURL = &s.config.DefaultProfilePictureURL
	}

	return customerDetails, nil
}
