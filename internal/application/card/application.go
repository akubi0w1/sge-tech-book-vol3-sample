package card

import (
	"context"
	"math/rand"

	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/domain/entity/shard"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/domain/repository/card"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/domain/repository/usercard"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/terror"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/util/idutil"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/util/timeutil"
)

type Application interface {
	DrawCard(ctx context.Context, userID string) (*shard.UserCard, error)
	ListCards(ctx context.Context, userID string) (shard.UserCardSlice, error)
}

type application struct {
	cardRepository     card.Repository
	userCardRepository usercard.Repository
}

func New(
	cardRepository card.Repository,
	userCardRepository usercard.Repository,
) Application {
	return &application{
		cardRepository:     cardRepository,
		userCardRepository: userCardRepository,
	}
}

func (a *application) DrawCard(ctx context.Context, userID string) (*shard.UserCard, error) {
	// 雑に抽選
	masterCards, err := a.cardRepository.SelectAll(ctx)
	if err != nil {
		return nil, err
	}

	drawIdx := rand.Intn(len(masterCards))
	drawCard := masterCards[drawIdx]

	now := timeutil.Now()
	id, err := idutil.UUID()
	if err != nil {
		return nil, err
	}

	userCard, err := a.userCardRepository.FindByCardID(ctx, userID, drawCard.ID)
	if err != nil {
		if terror.GetCode(err) != terror.CodeNotFound {
			return nil, err
		}

		// 新規作成
		userCard = &shard.UserCard{
			ID:        id,
			UserID:    userID,
			CardID:    drawCard.ID,
			Count:     1,
			CreatedAt: now,
			UpdatedAt: now,
		}

		err = a.userCardRepository.Insert(ctx, userCard)
		if err != nil {
			return nil, err
		}

		return userCard, nil
	}

	// 更新
	userCard.Count += 1
	userCard.UpdatedAt = now
	if err = a.userCardRepository.Update(ctx, userCard); err != nil {
		return nil, err
	}

	return userCard, nil
}

func (a *application) ListCards(ctx context.Context, userID string) (shard.UserCardSlice, error) {
	return a.userCardRepository.SelectByUserID(ctx, userID)
}
