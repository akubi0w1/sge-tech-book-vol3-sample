package card

import (
	"context"

	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/application/card"
	pb "github.com/akubi0w1/sge-tech-book-vol3-sample/pkg/pb/service"
)

type handler struct {
	cardApplication card.Application
}

func New(cardApplication card.Application) pb.CardServiceServer {
	return &handler{
		cardApplication: cardApplication,
	}
}

// DrawCard
func (h *handler) DrawCard(ctx context.Context, req *pb.DrawCardRequest) (*pb.DrawCardResponse, error) {
	card, err := h.cardApplication.DrawCard(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	return &pb.DrawCardResponse{
		Card: toCard(card),
	}, nil
}

// ListCard
func (h *handler) ListCard(ctx context.Context, req *pb.ListCardRequest) (*pb.ListCardResponse, error) {
	cards, err := h.cardApplication.ListCards(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}

	return &pb.ListCardResponse{
		Cards: toCardList(cards),
	}, nil
}
