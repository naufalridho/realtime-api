package _default

import (
	"context"

	utilctx "github.com/naufalridho/realtime-api/common/context"
	"github.com/naufalridho/realtime-api/model"
	"github.com/naufalridho/realtime-api/repository"
	"github.com/naufalridho/realtime-api/service"
)

type messageDefaultService struct {
	messageRepository repository.MessageRepository
}

func New(mr repository.MessageRepository) service.MessageService {
	return &messageDefaultService{
		messageRepository: mr,
	}
}

func (m *messageDefaultService) GetAll(ctx context.Context) []*model.Message {
	span, ctx := utilctx.StartSpanFromContext(ctx)
	defer span.Finish()

	return m.messageRepository.GetAll(ctx)
}

func (m *messageDefaultService) Send(ctx context.Context, message *model.Message) {
	span, ctx := utilctx.StartSpanFromContext(ctx)
	defer span.Finish()

	m.messageRepository.Create(ctx, message)
}

