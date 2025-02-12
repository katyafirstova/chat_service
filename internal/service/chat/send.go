package chat

import (
	"context"

	"github.com/katyafirstova/chat_service/internal/model"
)

func (s *serv) Send(ctx context.Context, req model.SendMessage) error {
	err := s.chatRepository.Send(ctx, req)
	if err != nil {
		return err
	}
	return nil
}
