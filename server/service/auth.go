package service

import (
	"context"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	pb "github.com/k197781/go-auth-proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	CreateToken(context.Context, *pb.CreateTokenRequest) (*pb.CreateTokenResponse, error)
}

type auth struct {
	secretKey string
}

func (a auth) CreateToken(ctx context.Context, in *pb.CreateTokenRequest) (*pb.CreateTokenResponse, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   in.Id,
		"name": in.Name,
		"nbf":  now.AddDate(0, 0, 7).Unix(),
	})
	tokenString, err := token.SignedString([]byte(a.secretKey))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to create jwt token.", err)
	}
	return &pb.CreateTokenResponse{
		Token: tokenString,
	}, nil
}

func NewAuthService(secretKey string) Auth {
	return auth{
		secretKey: secretKey,
	}
}
