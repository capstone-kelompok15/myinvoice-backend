package impl

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type notificationRepository struct {
	db  *sqlx.DB
	log *logrus.Entry
}

type NotificationRepositoryParams struct {
	DB  *sqlx.DB
	Log *logrus.Entry
}

func NewNotificationRepository(params *NotificationRepositoryParams) *notificationRepository {
	return &notificationRepository{
		db:  params.DB,
		log: params.Log,
	}
}
