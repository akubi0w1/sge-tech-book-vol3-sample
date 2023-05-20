package character

import (
	"context"

	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/domain/entity/master"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/domain/repository/character"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/infra/mysql"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/terror"
)

type repository struct {
	masterDB *mysql.MasterDB
}

func New(db *mysql.MasterDB) character.Repository {
	return &repository{
		masterDB: db,
	}
}

func (repo *repository) SelectAll(ctx context.Context) (master.CharacterSlice, error) {
	records, err := master.Characters().All(ctx, repo.masterDB)
	if err != nil {
		return nil, terror.Wrapf(terror.CodeInternal, err, "failed to get all character.")
	}

	return records, nil
}
