package api

import (
	"context"

	"github.com/katyafirstova/chat_service/internal/converter"
	"github.com/katyafirstova/chat_service/pkg/chat_v1"
)

func (i *Implementation) Create(ctx context.Context, req *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {
	uuid, err := i.chatService.Create(ctx, converter.CreateChatToServiceFromAPI(req))
	if err != nil {
		return nil, err
	}

	return &chat_v1.CreateResponse{Uuid: uuid}, nil
}
