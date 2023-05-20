package usercard

import (
	"context"

	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/domain/entity/shard"
)

type Repository interface {
	SelectByUserID(ctx context.Context, userID string) (shard.UserCardSlice, error)
	FindByCardID(ctx context.Context, userID string, cardID uint32) (*shard.UserCard, error)
	Insert(ctx context.Context, userCard *shard.UserCard) error
	Update(ctx context.Context, userCard *shard.UserCard) error
}
