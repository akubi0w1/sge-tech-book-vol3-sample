package user

import (
	"context"
	"database/sql"

	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/domain/entity/shard"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/domain/repository/user"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/infra/mysql"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/terror"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type repository struct {
	shardDB *mysql.ShardDB
}

func New(db *mysql.ShardDB) user.Repository {
	return &repository{
		shardDB: db,
	}
}

func (repo *repository) SelectAll(ctx context.Context) (shard.UserSlice, error) {
	records, err := shard.Users().All(ctx, repo.shardDB)
	if err != nil {
		return nil, terror.Wrapf(terror.CodeInternal, err, "failed to get all user.")
	}

	return records, nil
}

func (repo *repository) FindByID(ctx context.Context, userID string) (*shard.User, error) {
	user, err := shard.Users(
		shard.UserWhere.ID.EQ(userID),
	).One(ctx, repo.shardDB)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, terror.Wrapf(terror.CodeNotFound, err, "not found user. userId=%s", userID)
		}
		return nil, terror.Wrapf(terror.CodeInternal, err, "failed to find user. userId=%s", userID)
	}

	return user, nil
}

func (repo *repository) Insert(ctx context.Context, user *shard.User) error {
	if err := user.Insert(ctx, repo.shardDB, boil.Infer()); err != nil {
		return terror.Wrapf(terror.CodeInternal, err, "failed to insert user. user=%+v", user)
	}

	return nil
}

func (repo *repository) Update(ctx context.Context, user *shard.User) error {
	if _, err := user.Update(ctx, repo.shardDB, boil.Infer()); err != nil {
		return terror.Wrapf(terror.CodeInternal, err, "failed to update user. user=%+v", user)
	}

	return nil
}
