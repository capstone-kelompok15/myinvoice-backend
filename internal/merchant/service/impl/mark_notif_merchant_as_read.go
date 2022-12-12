package impl

import (
	"context"
)

func (s *merchantService) MarkNotifMerchantAsRead(ctx context.Context, NotifID int, MerchantID int) error {
	err := s.repoNotif.CheckNotifMerchantExist(NotifID)
	if err != nil {
		return err
	}
	return s.repoNotif.MarkNotifMerchantAsRead(ctx, NotifID, MerchantID)
}
