package impl

import (
	"context"
)

func (s *merchantService) GetUnreadNotifCount(ctx context.Context, merchantID int) (int, error) {

	count, err := s.repoNotif.GetUnreadNotifCountMerchant(ctx, merchantID)
	if err != nil {
		s.log.Warningln("[GetUnreadNotifCount] Failed on get repository", err.Error())
		return 0, err
	}

	return count, nil
}
