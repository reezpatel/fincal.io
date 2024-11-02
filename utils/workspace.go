package utils

import (
	"context"
	"errors"
	"strconv"

	"google.golang.org/grpc/metadata"
)

func GetWorkspaceId(ctx context.Context) (int64, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return 0, errors.New("error: failed to get metadata")
	}

	id, ok := md["x-workspace-id"]

	if !ok {
		return 0, errors.New("error: Failed to get workspace id")
	}

	return strconv.ParseInt(id[0], 10, 64)
}

type AuthAccess struct {
	U, W int64
}

type UserID struct{}

func AuthContext(ctx context.Context) *AuthAccess {
	w, err := GetWorkspaceId(ctx)

	if err != nil {
		return nil
	}

	userID, ok := ctx.Value(UserID{}).(int64)

	if !ok {
		return nil
	}

	return &AuthAccess{U: userID, W: w}
}
