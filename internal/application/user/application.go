package user

import (
	"context"

	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/domain/entity/shard"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/domain/repository/user"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/util/idutil"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/util/timeutil"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/pkg/pb/enums"
)

type Application interface {
	ListUsers(ctx context.Context) (shard.UserSlice, error)
	GetUser(ctx context.Context, userID string) (*shard.User, error)
	UpdateUserName(ctx context.Context, userID, name string) (*shard.User, error)
	CreateUser(ctx context.Context, name string, platform enums.Platform) (*shard.User, error)
}

type application struct {
	userRepository user.Repository
}

func New(userRepository user.Repository) Application {
	return &application{
		userRepository: userRepository,
	}
}

func (a *application) ListUsers(ctx context.Context) (shard.UserSlice, error) {
	return a.userRepository.SelectAll(ctx)
}

func (a *application) GetUser(ctx context.Context, userID string) (*shard.User, error) {
	return a.userRepository.FindByID(ctx, userID)
}

func (a *application) CreateUser(ctx context.Context, name string, platform enums.Platform) (*shard.User, error) {
	id, err := idutil.UUID()
	if err != nil {
		return nil, err
	}

	createdAt := timeutil.Now()
	user := &shard.User{
		ID:        id,
		Name:      name,
		Platform:  int32(platform),
		CreatedAt: createdAt,
		UpdatedAt: createdAt,
	}

	if err := a.userRepository.Insert(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (a *application) UpdateUserName(ctx context.Context, userID, name string) (*shard.User, error) {
	user, err := a.userRepository.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	user.Name = name
	user.UpdatedAt = timeutil.Now()

	if err := a.userRepository.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
