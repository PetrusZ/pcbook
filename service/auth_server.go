package service

import (
	"context"

	"github.com/PetrusZ/pcbook/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	userStore UserStore
	jwtManger *JWTManager
}

func NewAuthServer(userStore UserStore, jwtManger *JWTManager) *AuthServer {
	return &AuthServer{
		userStore: userStore,
		jwtManger: jwtManger,
	}
}

func (server *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := server.userStore.Find(req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "cannot find user: %v", user)
	}

	if user == nil || !user.IscorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username or password")
	}

	token, err := server.jwtManger.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token: %v", err)
	}

	res := &pb.LoginResponse{
		AccessToken: token,
	}
	return res, nil
}
