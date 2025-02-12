package app

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"

	chat "github.com/katyafirstova/chat_service/internal/api/chat"
	"github.com/katyafirstova/chat_service/internal/closer"
	"github.com/katyafirstova/chat_service/internal/config"
	"github.com/katyafirstova/chat_service/internal/config/env"
	"github.com/katyafirstova/chat_service/internal/repository"
	chatRepo "github.com/katyafirstova/chat_service/internal/repository/chat"
	"github.com/katyafirstova/chat_service/internal/service"
	chatService "github.com/katyafirstova/chat_service/internal/service/chat"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig
	pool       *pgxpool.Pool

	chatRepo    repository.ChatRepository
	chatService service.ChatService
	chatImpl    *chat.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) Pool(ctx context.Context) *pgxpool.Pool {
	if s.pool == nil {
		pool, err := pgxpool.Connect(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %s", err.Error())
		}
		closer.Add(func() error {
			pool.Close()
			return nil
		})

		s.pool = pool
	}

	return s.pool
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepo == nil {
		s.chatRepo = chatRepo.NewRepository(s.Pool(ctx))
	}

	return s.chatRepo
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = chatService.NewService(s.ChatRepository(ctx))
	}

	return s.chatService
}

func (s *serviceProvider) ChatImpl(ctx context.Context) *chat.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chat.NewImplementation(s.ChatService(ctx))
	}

	return s.chatImpl
}
