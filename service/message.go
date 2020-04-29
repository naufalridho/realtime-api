package service

import (
	"context"
	"github.com/naufalridho/realtime-api/model"
)

type MessageService interface {
	GetAll(ctx context.Context) []*model.Message
	Send(ctx context.Context, message *model.Message)
}