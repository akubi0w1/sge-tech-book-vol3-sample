package card

import (
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/domain/entity/shard"
	pbmodel "github.com/akubi0w1/sge-tech-book-vol3-sample/pkg/pb/model"
)

func toCardList(lst shard.UserCardSlice) []*pbmodel.Card {
	result := make([]*pbmodel.Card, 0, len(lst))
	for _, v := range lst {
		result = append(result, toCard(v))
	}
	return result
}

func toCard(v *shard.UserCard) *pbmodel.Card {
	return &pbmodel.Card{
		Id:        v.ID,
		UserId:    v.UserID,
		CardId:    v.CardID,
		CreatedAt: v.CreatedAt.Unix(),
		UpdatedAt: v.UpdatedAt.Unix(),
	}
}
