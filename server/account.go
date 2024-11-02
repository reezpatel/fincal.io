package server

import (
	"context"
	"fincal-server/fincal"
	"fincal-server/lib"
	"fincal-server/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"fincal-server/.gen/fincal/public/model"
	. "fincal-server/.gen/fincal/public/table"

	. "github.com/go-jet/jet/v2/postgres"
)

type AccountServer struct {
	fincal.UnimplementedAccountServiceServer
}

func (s *AccountServer) DepositAccounts(c context.Context, in *emptypb.Empty) (*fincal.DepositAccountsResponse, error) {
	auth := utils.AuthContext(c)

	if auth == nil {
		return nil, status.Errorf(codes.Unauthenticated, "AUTH_ERROR")
	}

	stmt := DepositAccount.SELECT(DepositAccount.AllColumns).WHERE(DepositAccount.WorkspaceID.EQ(Int64(auth.W)))

	var dest []struct {
		model.DepositAccount
	}

	err := stmt.Query(lib.DB, &dest)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}

	var accounts = make([]*fincal.DepositAccount, len(dest))

	for _, a := range dest {
		accounts = append(accounts, &fincal.DepositAccount{
			Id:             a.ID,
			WorkspaceId:    a.WorkspaceID,
			UserId:         a.UserID,
			CurrencyCode:   a.CurrencyCode,
			OpeningBalance: float64(a.OpeningBalance),
			Name:           *a.Name,
			Icon:           a.Icon,
			Type:           a.Type,
			Balance:        float64(a.Balance),
			CreatedAt:      timestamppb.New(*a.CreatedAt),
			Json:           *a.Json,
		})
	}

	return &fincal.DepositAccountsResponse{
		Accounts: accounts,
	}, nil
}

// func (s *AccountServer) CreateDepositAccount(context.Context, *fincal.CreateDepositAccountRequest) (*fincal.CreateDepositAccountResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method CreateDepositAccount not implemented")
// }
// func (s *AccountServer) DeleteDepositAccount(context.Context, *fincal.DeleteDepositAccountRequest) (*fincal.DeleteDepositAccountResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method DeleteDepositAccount not implemented")
// }
// func (s *AccountServer) CreditCardAccounts(context.Context, *fincal.emptypb.Empty) (*fincal.CreditCardAccountsResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method CreditCardAccounts not implemented")
// }
// func (s *AccountServer) CreateCreditCardAccount(context.Context, *fincal.CreateCreditCardAccountRequest) (*fincal.CreateCreditCardAccountResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method CreateCreditCardAccount not implemented")
// }
// func (s *AccountServer) DeleteCreditCardAccount(context.Context, *fincal.DeleteCreditCardAccountRequest) (*fincal.DeleteCreditCardAccountResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method DeleteCreditCardAccount not implemented")
// }
// func (s *AccountServer) SecurityAccounts(context.Context, *fincal.emptypb.Empty) (*fincal.SecurityAccountsResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method SecurityAccounts not implemented")
// }
// func (s *AccountServer) CreateSecurityAccount(context.Context, *fincal.CreateSecurityAccountRequest) (*fincal.CreateSecurityAccountResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method CreateSecurityAccount not implemented")
// }
// func (s *AccountServer) DeleteSecurityAccount(context.Context, *fincal.DeleteSecurityAccountRequest) (*fincal.DeleteSecurityAccountResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method DeleteSecurityAccount not implemented")
// }
// func (s *AccountServer) SecurityItems(context.Context, *emptypb.Empty) (*fincal.SecurityItemsResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method SecurityItems not implemented")
// }
// func (s *AccountServer) CreateSecurityItem(context.Context, *fincal.CreateSecurityItemRequest) (*fincal.CreateSecurityItemResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method CreateSecurityItem not implemented")
// }
// func (s *AccountServer) DeleteSecurityItem(context.Context, *fincal.DeleteSecurityItemRequest) (*fincal.DeleteSecurityItemResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method DeleteSecurityItem not implemented")
// }
