package user

import (
	"context"

	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/application/user"
	pb "github.com/akubi0w1/sge-tech-book-vol3-sample/pkg/pb/service"
)

type handler struct {
	userApplication user.Application
}

func New(userApplication user.Application) pb.UserServiceServer {
	return &handler{
		userApplication: userApplication,
	}
}

// ListUser
func (h *handler) ListUser(ctx context.Context, _ *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	users, err := h.userApplication.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.ListUserResponse{
		Users: toUserList(users),
	}, nil
}

// GetUser
func (h *handler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := h.userApplication.GetUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		User: toUser(user),
	}, nil
}

// CreateUser
func (h *handler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := h.userApplication.CreateUser(ctx, req.GetName(), req.GetPlatform())
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		User: toUser(user),
	}, nil
}

// UpdateUser
func (h *handler) UpdateUserName(ctx context.Context, req *pb.UpdateUserNameRequest) (*pb.UpdateUserNameResponse, error) {
	user, err := h.userApplication.UpdateUserName(ctx, req.GetUserId(), req.GetName())
	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserNameResponse{
		User: toUser(user),
	}, nil
}
