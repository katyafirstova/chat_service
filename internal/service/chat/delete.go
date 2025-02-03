package chat

import (
	"context"
)

func (s *serv) Delete(ctx context.Context, uuid string) error {
	err := s.chatRepository.Delete(ctx, uuid)
	if err != nil {
		return err
	}
	return nil
}
