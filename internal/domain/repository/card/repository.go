package card

import (
	"context"

	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/domain/entity/master"
)

type Repository interface {
	SelectAll(ctx context.Context) (master.CardSlice, error)
}
