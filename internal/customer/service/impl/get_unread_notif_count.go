package impl

import (
	"context"
)

func (s *customerService) GetUnreadNotifCount(ctx context.Context, customerID int) (int, error) {

	count, err := s.repoNotif.GetUnreadNotifCountCustomer(ctx, customerID)
	if err != nil {
		s.log.Warningln("[GetUnreadNotifCount] Failed on get repository", err.Error())
		return 0, err
	}

	return count, nil
}
