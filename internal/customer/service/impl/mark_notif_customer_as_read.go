package impl

import (
	"context"
)

func (s *customerService) MarkNotifCustomerAsRead(ctx context.Context, NotifID int, CustomerID int) error {
	err := s.repoNotif.CheckNotifCustomerExist(NotifID)
	if err != nil {
		return err
	}
	return s.repoNotif.MarkNotifCustomerAsRead(ctx, NotifID, CustomerID)
}
