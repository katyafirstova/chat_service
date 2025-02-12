package chat

import (
	"context"

	"github.com/katyafirstova/chat_service/internal/model"
)

func (s *serv) Create(ctx context.Context, req model.CreateChat) (string, error) {
	uuid, err := s.chatRepository.Create(ctx, req)
	if err != nil {
		return "", err
	}

	return uuid, nil
}
