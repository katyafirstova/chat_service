package api

import (
	"github.com/katyafirstova/chat_service/internal/service"
	"github.com/katyafirstova/chat_service/pkg/chat_v1"
)

type Implementation struct {
	chat_v1.UnimplementedChatV1Server
	chatService service.ChatService
}

func NewImplementation(chatService service.ChatService) *Implementation {
	return &Implementation{chatService: chatService}
}
