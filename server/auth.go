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
)

type AuthServer struct {
	fincal.UnimplementedAuthServiceServer
}

func (s *AuthServer) Register(ctx context.Context, in *fincal.RegisterAccountRequest) (*fincal.RegisterAccountResponse, error) {
	err := protovalidate.Validate(in)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}

	hex, err := utils.RandomHex(32)

	if err != nil {
		return nil, err
	}

	smt := User.INSERT(User.MutableColumns).MODEL(&model.User{
		Email:    in.Email,
		Icon:     "",
		Password: &[]string{utils.Hash(fmt.Sprintf("%s%s", in.Password, hex))}[0],
		Name:     &[]string{in.Name}[0],
		Hash:     hex,
	}).RETURNING(User.ID)

	var dest []struct {
		model.User
	}

	err = smt.Query(lib.DB, &dest)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}

	jwt, err := lib.CreateJWTToken(dest[0].ID)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}

	return &fincal.RegisterAccountResponse{
		Success: true,
		Error:   "",
		Token:   jwt,
	}, nil
}

func (s *AuthServer) SignIn(ctx context.Context, in *fincal.SignInRequest) (*fincal.SignInResponse, error) {
	err := protovalidate.Validate(in)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}

	stmt := User.SELECT(User.Email, User.Hash).WHERE(User.Email.EQ(String(in.Email)))

	var dest []struct {
		model.User
	}

	err = stmt.Query(lib.DB, &dest)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}

	if len(dest) == 0 {
		return &fincal.SignInResponse{
			Success: false,
			Error:   "INVALID_CREDENTIALS",
		}, nil
	}

	pwd := utils.Hash(fmt.Sprintf("%s%s", in.Password, dest[0].Hash))

	smt2 := User.SELECT(User.ID).WHERE(AND(User.Email.EQ(String(in.Email)), User.Password.EQ(String(pwd))))

	var dest2 []struct {
		model.User
	}

	err = smt2.Query(lib.DB, &dest2)

	if err != nil {
		return nil, err
	}

	if len(dest2) == 0 {
		return &fincal.SignInResponse{
			Success: false,
			Error:   "INVALID_CREDENTIALS",
		}, nil
	}

	jwt, err := lib.CreateJWTToken(dest2[0].ID)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}

	return &fincal.SignInResponse{
		Success: true,
		Error:   "",
		Token:   jwt,
	}, nil
}
