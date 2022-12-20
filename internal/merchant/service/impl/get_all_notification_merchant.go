package impl

import (
	"context"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	notifUtils "github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/notifications"
)

func (s *merchantService) GetAllNotificationMerchant(ctx context.Context, merchantID int, req *dto.NotificationRequest) (*[]dto.NotificationRespond, error) {
	notifications, err := s.repoNotif.GetAllNotificationMerchant(ctx, merchantID, req)
	if err != nil {

		s.log.Warningln("[GetAllNotificationMerchant] Error while getting all notification :", err.Error())
		return nil, err
	}

	respond := []dto.NotificationRespond{}

	for i := 0; i < len(*notifications); i++ {
		respond = append(respond, dto.NotificationRespond{
			ID:               (*notifications)[i].ID,
			InvoiceID:        (*notifications)[i].InvoiceID,
			NotificationType: (*notifications)[i].Type,
			Title:            (*notifications)[i].Title,
			IsRead:           (*notifications)[i].IsRead,
			Content:          notifUtils.GenerateContent((*notifications)[i].Title, (*notifications)[i].CustomerName, (*notifications)[i].InvoiceID),
			CreatedAt:        (*notifications)[i].CreatedAt,
		})

	}

	return &respond, nil
}
