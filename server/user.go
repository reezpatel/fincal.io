package server

import (
	"context"
	"fincal-server/.gen/fincal/public/model"
	. "fincal-server/.gen/fincal/public/table"
	"fincal-server/fincal"
	"fincal-server/lib"
	"fincal-server/utils"
	"fmt"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/go-jet/jet/v2/postgres"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserServer struct {
	fincal.UnimplementedUserServiceServer
}

func (s *UserServer) Self(ctx context.Context, in *emptypb.Empty) (*fincal.UserResponse, error) {
	userID, ok := ctx.Value(utils.UserID{}).(int64)

	if !ok {
		return nil, status.Error(codes.Internal, "Error: Failed to get user id")
	}

	stmt := User.SELECT(User.Email, User.Name, User.Icon).WHERE(User.ID.EQ(Int64(userID)))

	var dest []struct {
		model.User
	}

	err := stmt.Query(lib.DB, &dest)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}

	if len(dest) == 0 {
		return nil, status.Error(codes.Internal, "Error: User not found")
	}

	return &fincal.UserResponse{User: &fincal.User{
		Id:    dest[0].ID,
		Email: dest[0].Email,
		Icon:  dest[0].Icon,
		Name:  *dest[0].Name,
	}}, nil
}

func (s *UserServer) UpdateSelf(ctx context.Context, in *fincal.UpdateSelfRequest) (*fincal.UpdateSelfResponse, error) {
	err := protovalidate.Validate(in)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}

	userID, ok := ctx.Value(utils.UserID{}).(int64)

	if !ok {
		return nil, status.Error(codes.Internal, "Error: Failed to get user id")
	}

	stmt := User.UPDATE(User.Name, User.Icon).WHERE(User.ID.EQ(Int64(userID))).SET(in.Name, in.Icon)

	_, err = stmt.Exec(lib.DB)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}

	return &fincal.UpdateSelfResponse{
		Success: true,
		Error:   "",
	}, nil
}

func (s *UserServer) UpdatePassword(ctx context.Context, in *fincal.UpdatePasswordRequest) (*fincal.UpdatePasswordResponse, error) {
	err := protovalidate.Validate(in)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}

	userID, ok := ctx.Value(utils.UserID{}).(int64)

	if !ok {
		return nil, status.Error(codes.Internal, "Error: Failed to get user id")
	}

	stmt := User.SELECT(User.Password, User.Hash).WHERE(User.ID.EQ(Int64(userID)))

	var dest []struct {
		model.User
	}

	err = stmt.Query(lib.DB, &dest)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}

	if len(dest) == 0 {
		return nil, status.Error(codes.Internal, "Error: User not found")
	}

	hash1 := utils.Hash(fmt.Sprintf("%s%s", in.CurrentPassword, dest[0].Hash))
	hash2 := utils.Hash(fmt.Sprintf("%s%s", dest[0].Password, dest[0].Hash))

	if hash1 != hash2 {
		return &fincal.UpdatePasswordResponse{
			Success: false,
			Error:   "INVALID_PASSWORD",
		}, nil
	}

	us := User.UPDATE(User.Password).WHERE(User.ID.EQ(Int64(userID))).SET(in.NewPassword)

	_, err = us.Exec(lib.DB)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}

	return &fincal.UpdatePasswordResponse{
		Success: true,
		Error:   "",
	}, nil
}
