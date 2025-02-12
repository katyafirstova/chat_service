package repository

import (
	"context"

	"github.com/katyafirstova/chat_service/internal/model"
)

type ChatRepository interface {
	Create(ctx context.Context, req model.CreateChat) (string, error)
	Send(ctx context.Context, req model.SendMessage) error
	Delete(ctx context.Context, uuid string) error
}
