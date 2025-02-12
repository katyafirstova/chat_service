package chat

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/katyafirstova/chat_service/internal/model"
	def "github.com/katyafirstova/chat_service/internal/repository"
)

var _ def.ChatRepository = (*repo)(nil)

const (
	chatTable               = "chats"
	chatTableColumnUUID     = "chat_uuid"
	chatTableColumnUserUUID = "user_uuid"

	messageTable               = "messages"
	messageTableColumnUUID     = "uuid"
	messageTableColumnUserUUID = "user_uuid"
	messageTableColumnChatUUID = "chat_uuid"
	messageTableColumnText     = "text"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *repo {
	return &repo{db: db}
}

func (r *repo) Create(ctx context.Context, req model.CreateChat) (string, error) {
	chatUUID := uuid.NewString()

	builderInsert := sq.Insert(chatTable).
		PlaceholderFormat(sq.Dollar).
		Columns(chatTableColumnUUID, chatTableColumnUserUUID)

	for _, userUUID := range req.UserUuids {
		builderInsert = builderInsert.Values(chatUUID, userUUID)
	}

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return "", fmt.Errorf("failed to build SQL query: %w", err)
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return "", err
	}

	return chatUUID, nil
}

func (r *repo) Send(ctx context.Context, req model.SendMessage) error {
	builderInsert := sq.Insert(messageTable).
		PlaceholderFormat(sq.Dollar).
		Columns(messageTableColumnUUID, messageTableColumnUserUUID, messageTableColumnChatUUID, messageTableColumnText).
		Values(uuid.NewString(), req.UserUuid, req.ChatUuid, req.Text)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Delete(ctx context.Context, uuid string) error {
	builderDelete := sq.Delete(messageTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{messageTableColumnChatUUID: uuid})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	builderDelete = sq.Delete(chatTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{chatTableColumnUUID: uuid})

	query, args, err = builderDelete.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
