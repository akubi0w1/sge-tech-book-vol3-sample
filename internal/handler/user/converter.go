package user

import (
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/domain/entity/shard"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/pkg/pb/enums"
	pbmodel "github.com/akubi0w1/sge-tech-book-vol3-sample/pkg/pb/model"
)

func toUserList(lst shard.UserSlice) []*pbmodel.User {
	result := make([]*pbmodel.User, 0, len(lst))
	for _, v := range lst {
		result = append(result, toUser(v))
	}
	return result
}

func toUser(v *shard.User) *pbmodel.User {
	return &pbmodel.User{
		Id:        v.ID,
		Name:      v.Name,
		Platform:  enums.Platform(v.Platform),
		CreatedAt: v.CreatedAt.Unix(),
		UpdatedAt: v.UpdatedAt.Unix(),
	}
}
