package impl

import (
	"context"
	"fmt"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/go-redis/redis/v8"
)

func (s *authService) CustomerEmailVerification(ctx context.Context, req *dto.CustomerEmailVerification) error {
	redisKey := fmt.Sprintf("customer-regis:%s", req.Email)

	code, err := s.redis.Get(ctx, redisKey).Result()

	if err == redis.Nil {
		return customerrors.ErrUnauthorized
	}

	if err != nil {
		s.log.Warningln("[CustomerEmailVerification] Error while getting the code on redis", err.Error())
		return err
	}

	if code != req.Code {
		return customerrors.ErrUnauthorized
	}

	err = s.repo.CustomerEmailVerification(ctx, req)
	if err != nil {
		s.log.Warningln("[CustomerEmailVerification] Error while updating the is_verified on customer", err.Error())
		return err
	}

	err = s.redis.Del(ctx, redisKey).Err()
	if err != nil {
		s.log.Warningln("[CustomerEmailVerification] Error while destroying the redis key", err.Error())
		return err
	}

	return nil
}
