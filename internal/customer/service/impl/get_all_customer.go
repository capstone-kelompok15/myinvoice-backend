package impl

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (s *customerService) GetAllCustomer(ctx context.Context, req *dto.GetAllCustomerRequest) (*[]dto.GetAllCustomerRespond, error) {
	customers, err := s.repo.GetAllCustomer(ctx, req)
	if err != nil {

		s.log.Warningln("[GetAllCustomers] Error while getting all customer :", err.Error())
		return nil, err
	}

	for i := 0; i < len(*customers); i++ {
		if (*customers)[i].DisplayProfileURL == nil {
			(*customers)[i].DisplayProfileURL = &s.config.DefaultProfilePictureURL
		}
	}

	return customers, nil
}
