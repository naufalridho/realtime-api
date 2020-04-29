package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"

	utilctx "github.com/naufalridho/realtime-api/common/context"
	"github.com/naufalridho/realtime-api/model"
	"github.com/naufalridho/realtime-api/service"
)

type SendRequest struct {
	Message model.Message `json:"message"`
}

type Response struct {
	Success	bool   		`json:"success"`
	ErrMsg	string 		`json:"errorMsg"`
	Data	interface{} `json:"data,omitempty"`
}

type MessageHttpHandlerV1 struct {
	messageService service.MessageService
}

func New(ms service.MessageService) *MessageHttpHandlerV1 {
	return &MessageHttpHandlerV1{
		messageService: ms,
	}
}

func (m *MessageHttpHandlerV1) SendMessage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := utilctx.NewContextFromRequest(r)

	statusCode := http.StatusInternalServerError
	response := &Response{
		Success: false,
		ErrMsg: "failed to send message",
	}

	defer func() {
		bytes, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write(bytes)
	}()

	req := SendRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		statusCode = http.StatusBadRequest
		return
	}

	if err := m.validateSendRequest(ctx, req); err != nil {
		statusCode = http.StatusBadRequest
		response.ErrMsg = err.Error()
		return
	}

	m.messageService.Send(ctx, &req.Message)
	statusCode = http.StatusOK
	response.Success = true
	response.ErrMsg = ""
}

func (m *MessageHttpHandlerV1) GetAllMessages(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ctx := utilctx.NewContextFromRequest(r)

	messages := m.messageService.GetAll(ctx)

	bytes, _ := json.Marshal(
		&Response{
			Success: true,
			Data: messages,
		})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func (m *MessageHttpHandlerV1) validateSendRequest(ctx context.Context, req SendRequest) error {
	span, ctx := utilctx.StartSpanFromContext(ctx)
	defer span.Finish()

	if req.Message.Content == "" {
		return fmt.Errorf("message content cannot be empty")
	}

	return nil
}
