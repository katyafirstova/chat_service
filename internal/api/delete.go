package api

import (
	"context"

	"github.com/katyafirstova/chat_service/pkg/chat_v1"
)

func (i *Implementation) Delete(ctx context.Context, req *chat_v1.DeleteRequest) error {
	err := i.chatService.Delete(ctx, req.Uuid)
	if err != nil {
		return err
	}

	return nil
}
