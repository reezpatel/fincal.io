package server

import (
	"context"
	"fincal-server/lib"
	"fincal-server/utils"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "AUTH_ERROR")
	}

	authHeaders := md["authorization"]
	if len(authHeaders) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "AUTH_ERROR")
	}

	token := authHeaders[0]

	if !strings.HasPrefix(token, "Bearer ") {
		return nil, status.Errorf(codes.Unauthenticated, "AUTH_ERROR")
	}

	jwtToken := strings.TrimPrefix(token, "Bearer ")

	userID, err := lib.ValidateJWTToken(jwtToken)

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "AUTH_ERROR")
	}

	newCtx := context.WithValue(ctx, utils.UserID{}, userID)

	return handler(newCtx, req)
}
