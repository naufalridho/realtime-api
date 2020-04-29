package repository

import (
	"context"
	"github.com/naufalridho/realtime-api/model"
)

type MessageRepository interface {
	GetAll(ctx context.Context) []*model.Message
	Create(ctx context.Context, message *model.Message)
}
