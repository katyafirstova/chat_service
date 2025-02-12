package api

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/katyafirstova/chat_service/internal/converter"
	"github.com/katyafirstova/chat_service/pkg/chat_v1"
)

func (i *Implementation) Send(ctx context.Context, req *chat_v1.SendRequest) (*emptypb.Empty, error) {
	err := i.chatService.Send(ctx, converter.SendMessageToServiceFromApi(req))
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
