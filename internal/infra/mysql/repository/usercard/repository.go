package usercard

import (
	"context"
	"database/sql"

	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/domain/entity/shard"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/domain/repository/usercard"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/infra/mysql"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/terror"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type repository struct {
	shardDB *mysql.ShardDB
}

func New(db *mysql.ShardDB) usercard.Repository {
	return &repository{
		shardDB: db,
	}
}

func (repo *repository) SelectByUserID(ctx context.Context, userID string) (shard.UserCardSlice, error) {
	records, err := shard.UserCards(shard.UserCardWhere.UserID.EQ(userID)).All(ctx, repo.shardDB)
	if err != nil {
		return nil, terror.Wrapf(terror.CodeInternal, err, "failed to get userCards by userId. userId=%s", userID)
	}

	return records, nil
}

func (repo *repository) FindByCardID(ctx context.Context, userID string, cardID uint32) (*shard.UserCard, error) {
	record, err := shard.UserCards(
		shard.UserCardWhere.UserID.EQ(userID),
		shard.UserCardWhere.CardID.EQ(cardID),
	).One(ctx, repo.shardDB)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, terror.Wrapf(terror.CodeNotFound, err, "failed to find userCard. userId=%s, cardId=%d", userID, cardID)
		}

		return nil, terror.Wrapf(terror.CodeInternal, err, "failed to get userCards. userId=%s, cardId=%d", userID, cardID)
	}

	return record, nil
}

func (repo *repository) Insert(ctx context.Context, userCard *shard.UserCard) error {
	if err := userCard.Insert(ctx, repo.shardDB, boil.Infer()); err != nil {
		return terror.Wrapf(terror.CodeInternal, err, "failed to insert userCard. userCard=%+v", userCard)
	}

	return nil
}

func (repo *repository) Update(ctx context.Context, userCard *shard.UserCard) error {
	if _, err := userCard.Update(ctx, repo.shardDB, boil.Infer()); err != nil {
		return terror.Wrapf(terror.CodeInternal, err, "failed to update userCard. userCard=%+v", userCard)
	}

	return nil
}
