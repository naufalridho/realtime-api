package inmemory

import (
	"context"

	utilctx "github.com/naufalridho/realtime-api/common/context"
	"github.com/naufalridho/realtime-api/model"
	"github.com/naufalridho/realtime-api/repository"
)

type messageInMemoryRepository struct {
	messages []*model.Message
}

func New() repository.MessageRepository {
	return &messageInMemoryRepository{}
}

func (m *messageInMemoryRepository) GetAll(ctx context.Context) []*model.Message {
	span, ctx := utilctx.StartSpanFromContext(ctx)
	defer span.Finish()

	return m.messages
}

func (m *messageInMemoryRepository) Create(ctx context.Context, message *model.Message) {
	span, ctx := utilctx.StartSpanFromContext(ctx)
	defer span.Finish()

	m.messages = append(m.messages, message)
}
