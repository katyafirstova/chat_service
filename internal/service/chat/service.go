package chat

import (
	"github.com/katyafirstova/chat_service/internal/repository"
)

type serv struct {
	chatRepository repository.ChatRepository
}

func NewService(
	chatRepository repository.ChatRepository,
) *serv {
	return &serv{
		chatRepository: chatRepository,
	}
}
