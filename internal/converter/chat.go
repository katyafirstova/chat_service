package converter

import (
	"github.com/katyafirstova/chat_service/internal/model"
	"github.com/katyafirstova/chat_service/pkg/chat_v1"
)

func CreateChatToServiceFromApi(req *chat_v1.CreateRequest) model.CreateChat {
	return model.CreateChat{
		UserUuids: req.UserUuids,
	}
}

func SendMessageToServiceFromApi(req *chat_v1.SendRequest) model.SendMessage {
	return model.SendMessage{
		UserUuid: req.SenderUuid,
		ChatUuid: req.ChatUuid,
		Text:     req.Text,
	}
}
