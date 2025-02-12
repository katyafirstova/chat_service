package api

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/katyafirstova/chat_service/pkg/chat_v1"
)

func (i *Implementation) Delete(ctx context.Context, req *chat_v1.DeleteRequest) (*emptypb.Empty, error) {
	err := i.chatService.Delete(ctx, req.Uuid)
	if err != nil {
		return &emptypb.Empty{}, err
	}

	return &emptypb.Empty{}, nil
}
