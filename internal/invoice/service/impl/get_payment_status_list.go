package impl

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (s *invoiceService) GetPaymentStatusList(ctx context.Context) (*[]dto.PaymentStatus, error) {
	paymentStatus, err := s.repo.GetPaymentStatusList(ctx)
	if err != nil {
		s.log.Warningln("[GetPaymentStatusList] Error on running the repository")
		return nil, err
	}
	return paymentStatus, nil
}
