package main

import (
	"context"
	"fmt"
	"log"
	"net"

	sq "github.com/Masterminds/squirrel"
	`github.com/google/uuid`
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/katyafirstova/chat_service/pkg/chat_v1"
)

const (
	dbDSN = "host=localhost port=54322 dbname=chat_db user=chat_user password=chat_password sslmode=disable"

	address                     = "127.0.0.1:50002"
	chatTable                   = "chats"
	messageTable                = "messages"
	chatTableColumnUUID         = "uuid"
	chatTableColumnUserUUID     = "user_uuid"
	chatTableColumnTimestamp    = "timestamp"
	messageTableColumnUUID      = "uuid"
	messageTableColumnUserUUID  = "user_uuid"
	messageTableColumnChatUUID  = "chat_uuid"
	messageTableColumnText      = "text"
	messageTableColumnCreatedAt = "created_at"
)

var pool *pgxpool.Pool

type server struct {
	chat_v1.UnimplementedChatV1Server
}

func (s *server) Create(ctx context.Context, req *chat_v1.CreateRequest) (*chat_v1.CreateResponse, error) {

	builderInsert := sq.Insert(chatTable).
		PlaceholderFormat(sq.Dollar).
		Columns(chatTableColumnUUID, chatTableColumnUserUUID, chatTableColumnTimestamp).
		Values(uuid.NewString(), req).
		Suffix(fmt.Sprintf("RETURNING %s", chatTableColumnUUID))

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return nil, err
	}

	var newUUID string
	err = pool.QueryRow(ctx, query, args...).Scan(&newUUID)
	if err != nil {
		return nil, err
	}

	return &chat_v1.CreateResponse{Uuid: newUUID}, nil

}

func (s *server) Delete(ctx context.Context, req *chat_v1.DeleteRequest) (*emptypb.Empty, error) {
	builderDelete := sq.Delete(chatTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{chatTableColumnUUID: req.Uuid})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *server) Send(ctx context.Context, req *chat_v1.SendRequest) (*emptypb.Empty, error) {
	builderInsert := sq.Insert(messageTable).
		PlaceholderFormat(sq.Dollar).
		Columns(messageTableColumnUUID, messageTableColumnUserUUID, messageTableColumnChatUUID, messageTableColumnText,
			messageTableColumnCreatedAt).
		Values(uuid.NewString(), req.SenderUuid, req.ChatUuid, req.Text)

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return nil, err
	}

	_, err = pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func main() {
	var err error
	ctx := context.Background()

	pool, err = pgxpool.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
	}
	defer pool.Close()

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to create listener: %s", err.Error())
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	chat_v1.RegisterChatV1Server(grpcServer, &server{})

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %s", err.Error())
	}
}
