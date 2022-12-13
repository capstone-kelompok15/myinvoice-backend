package impl

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (s *invoiceService) GetCustomers(ctx context.Context, req *dto.GetMerchantCustomerList) (*[]dto.BriefCustomer, int, error) {
	customers, count, err := s.repo.GetCustomers(ctx, req)
	if err != nil {
		s.log.Warningln("[GetCustomers] Error on running the service:", err.Error())
		return nil, 0, err
	}
	return customers, count, nil
}
