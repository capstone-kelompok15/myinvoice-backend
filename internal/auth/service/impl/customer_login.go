package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/capstone-kelompok15/myinvoice-backend/pkg/dto"
	customerrors "github.com/capstone-kelompok15/myinvoice-backend/pkg/errors"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/passwordutils"
	"github.com/capstone-kelompok15/myinvoice-backend/pkg/utils/tokenutils"
	"github.com/go-redis/redis/v8"
)

func (s *authService) CustomerLogin(ctx context.Context, req *dto.CustomerLoginRequest) (*dto.CustomerAccessToken, error) {
	req.Password = passwordutils.HashPassword(req.Password)

	customerContext, err := s.repo.AuthorizeCustomerLogin(ctx, req)
	if err != nil {
		s.log.Warningln("[CustomerLogin] Error while authorize customer login", err.Error())
		return nil, customerrors.ErrUnauthorized
	}
	customerContext.DeviceID = req.DeviceID

	err = s.repo.InvalidCustomerAccessToken(ctx, customerContext.ID)
	if err != nil {
		s.log.Warningln("[CustomerLogin] Error while calling the repo function", err.Error())
		return nil, err
	}

	deviceID := passwordutils.HashPassword(req.DeviceID)
	err = s.repo.InsertCustomerAccessToken(ctx, customerContext.ID, deviceID)
	if err != nil {
		s.log.Warningln("[CustomerLogin] Error while calling the repo function", err.Error())
		return nil, err
	}

	accessTokenStr, err := tokenutils.GenerateCustomerAccessToken(&dto.CustomerAccessTokenParams{
		DeviceInformation: req.DeviceID,
		UserInformation:   customerContext,
		Config:            &s.config.CustomerToken,
	})
	if err != nil {
		s.log.Warningln("[CustomerLogin] Error while generating customer access token", err.Error())
		return nil, err
	}

	accessToken := dto.CustomerAccessToken{
		AccessToken: *accessTokenStr,
	}

	customerContextJson, err := json.Marshal(*customerContext)
	if err != nil {
		s.log.Warningln("[CustomerLogin] Error while marshalling customer context", err.Error())
		return nil, err
	}

	redisKey := fmt.Sprintf("customer-access-token:%d", customerContext.ID)
	res := s.redis.Get(ctx, redisKey)
	if res.Err() != nil {
		if res.Err() == redis.Nil {
			s.redis.Set(ctx, redisKey, string(customerContextJson), 1*time.Hour)
		}
		s.log.Warningln("[CustomerLogin] Error while getting the cache from redis", err.Error())
		return nil, err
	} else if res.Val() != string(customerContextJson) {
		s.redis.Set(ctx, redisKey, string(customerContextJson), 1*time.Hour)
	}

	return &accessToken, nil
}
