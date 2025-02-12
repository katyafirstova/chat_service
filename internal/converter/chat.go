package converter

import (
	"github.com/katyafirstova/chat_service/internal/model"
	"github.com/katyafirstova/chat_service/pkg/chat_v1"
)

func CreateChatToServiceFromAPI(req *chat_v1.CreateRequest) model.CreateChat {
	return model.CreateChat{
		UserUuids: req.UserUuids,
	}
}

func SendMessageToServiceFromAPI(req *chat_v1.SendRequest) model.SendMessage {
	return model.SendMessage{
		UserUUID: req.SenderUuid,
		ChatUUID: req.ChatUuid,
		Text:     req.Text,
	}
}
