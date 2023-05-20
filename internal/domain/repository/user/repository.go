package user

import (
	"context"

	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/domain/entity/shard"
)

type Repository interface {
	SelectAll(ctx context.Context) (shard.UserSlice, error)
	FindByID(ctx context.Context, userID string) (*shard.User, error)
	Insert(ctx context.Context, user *shard.User) error
	Update(ctx context.Context, user *shard.User) error
}
