package impl

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
)

func (r *invoiceRepository) GetPaymentStatusList(ctx context.Context) (*[]dto.PaymentStatus, error) {
	getPaymentStatusSQL, _, err := squirrel.
		Select("id, status_name").
		From("payment_statuses").
		ToSql()
	if err != nil {
		r.log.Warningln("[GetPaymentStatusList] Error on build sql:", err.Error())
		return nil, err
	}

	paymentStatus := []dto.PaymentStatus{}
	err = r.db.SelectContext(ctx, &paymentStatus, getPaymentStatusSQL)
	if err != nil {
		r.log.Warningln("[GetPaymentStatusList] Error on execute the query:", err.Error())
		return nil, err
	}

	return &paymentStatus, nil
}
