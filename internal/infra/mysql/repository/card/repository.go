package card

import (
	"context"

	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/domain/entity/master"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/domain/repository/card"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/infra/mysql"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/terror"
)

type repository struct {
	masterDB *mysql.MasterDB
}

func New(db *mysql.MasterDB) card.Repository {
	return &repository{
		masterDB: db,
	}
}

func (repo *repository) SelectAll(ctx context.Context) (master.CardSlice, error) {
	records, err := master.Cards().All(ctx, repo.masterDB)
	if err != nil {
		return nil, terror.Wrapf(terror.CodeInternal, err, "failed to get all card.")
	}

	return records, nil
}
