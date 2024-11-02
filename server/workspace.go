package server

import (
	"context"
	"fincal-server/.gen/fincal/public/model"
	. "fincal-server/.gen/fincal/public/table"
	"fincal-server/fincal"
	"fincal-server/lib"
	"fincal-server/utils"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/go-jet/jet/v2/postgres"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type WorkspaceServer struct {
	fincal.UnimplementedWorkspaceServiceServer
}

func (s *WorkspaceServer) UpdateWorkspace(c context.Context, in *fincal.UpdateWorkspaceRequest) (*fincal.UpdateWorkspaceResponse, error) {
	userID, ok := c.Value(utils.UserID{}).(int64)

	if !ok {
		return nil, status.Error(codes.Internal, "Error: Failed to get user id")
	}

	if in.Id != 0 {
		role := lib.GetWorkspaceRole(&utils.AuthAccess{U: userID, W: in.Id})

		if role != model.Role_Owner {
			return nil, status.Errorf(codes.PermissionDenied, "NOT_OWNER")
		}

		smt := Workspace.UPDATE(Workspace.Name, Workspace.CurrencyCode, Workspace.Icon).SET(in.Name, in.CurrencyCode, in.Icon).WHERE(Workspace.ID.EQ(Int64(in.Id)))

		_, err := smt.Exec(lib.DB)

		if err != nil {
			return nil, status.Errorf(codes.Internal, "Error: %v", err)
		}

		return &fincal.UpdateWorkspaceResponse{
			Success: true,
			Error:   "",
		}, nil
	}

	role := lib.GetWorkspaceRole(&utils.AuthAccess{U: userID, W: in.Id})

	if role != model.Role_Owner {
		return nil, status.Errorf(codes.PermissionDenied, "NOT_OWNER")
	}

	smt := Workspace.INSERT(Workspace.MutableColumns).MODEL(&model.Workspace{
		Name:         in.Name,
		Icon:         in.Icon,
		CurrencyCode: in.CurrencyCode,
	}).RETURNING(Workspace.ID)

	var dest []struct {
		model.Workspace
	}

	err := smt.Query(lib.DB, &dest)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: Failed to get user id")
	}

	if len(dest) == 0 {
		return nil, status.Errorf(codes.Internal, "Error: Failed to get user id")
	}

	workspaceID := dest[0].ID

	smt2 := WorkspaceAccess.INSERT(WorkspaceAccess.MutableColumns).MODEL(&model.WorkspaceAccess{
		WorkspaceID: workspaceID,
		UserID:      userID,
		Role:        model.Role_Owner,
	})

	_, err = smt2.Exec(lib.DB)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: Failed to get user id")
	}

	lib.SetWorkspaceRole(utils.AuthAccess{U: userID, W: workspaceID}, model.Role_Owner)

	return &fincal.UpdateWorkspaceResponse{
		Success: true,
		Error:   "",
	}, nil
}

func (s *WorkspaceServer) Workspaces(c context.Context, in *emptypb.Empty) (*fincal.WorkspacesResponse, error) {
	userID, ok := c.Value(utils.UserID{}).(int64)

	if !ok {
		return nil, status.Error(codes.Internal, "Error: Failed to get user id")
	}

	stmt := WorkspaceAccess.INNER_JOIN(
		Workspace, Workspace.ID.EQ(WorkspaceAccess.WorkspaceID),
	).SELECT(Workspace.AllColumns).WHERE(WorkspaceAccess.UserID.EQ(Int64(userID))).ORDER_BY(Workspace.CreatedAt.DESC())

	var dest []struct {
		model.Workspace
	}

	err := stmt.Query(lib.DB, &dest)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %s", err)
	}

	var workspaces = make([]*fincal.Workspace, len(dest))

	for _, w := range dest {
		workspaces = append(workspaces, &fincal.Workspace{
			Id:           w.ID,
			Name:         w.Name,
			CurrencyCode: w.CurrencyCode,
			Icon:         w.Icon,
			CreatedAt:    timestamppb.New(*w.CreatedAt),
		})
	}

	return &fincal.WorkspacesResponse{
		Items: workspaces,
	}, nil
}

func (s *WorkspaceServer) DeleteWorkspace(c context.Context, in *fincal.DeleteWorkspaceRequest) (*fincal.DeleteWorkspaceResponse, error) {
	userID, ok := c.Value(utils.UserID{}).(int64)

	if !ok {
		return nil, status.Error(codes.Internal, "Error: Failed to get user id")
	}

	role := lib.GetWorkspaceRole(&utils.AuthAccess{U: userID, W: in.Id})

	if role != model.Role_Owner {
		return nil, status.Errorf(codes.PermissionDenied, "NOT_OWNER")
	}

	smt := Workspace.DELETE().WHERE(Workspace.ID.EQ(Int64(in.Id)))

	_, err := smt.Exec(lib.DB)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}

	return &fincal.DeleteWorkspaceResponse{
		Success: true,
		Error:   "",
	}, nil
}

func (s *WorkspaceServer) AddUserToWorkspace(c context.Context, in *fincal.AddUserToWorkspaceRequest) (*fincal.AddUserToWorkspaceResponse, error) {
	err := protovalidate.Validate(in)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}

	var role model.Role

	err = role.Scan(in.Access.Role)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}

	userID, ok := c.Value(utils.UserID{}).(int64)

	if !ok {
		return nil, status.Error(codes.Internal, "Error: Failed to get user id")
	}

	if in.Access.UserId == userID {
		return nil, status.Errorf(codes.PermissionDenied, "USER_CANNOT_MANAGE_SELF")
	}

	userRole := lib.GetWorkspaceRole(&utils.AuthAccess{U: userID, W: in.Access.WorkspaceId})

	if userRole != model.Role_Owner {
		return nil, status.Errorf(codes.PermissionDenied, "NOT_OWNER")
	}

	smt := WorkspaceAccess.INSERT(WorkspaceAccess.MutableColumns).MODEL(&model.WorkspaceAccess{
		WorkspaceID: in.Access.WorkspaceId,
		UserID:      in.Access.UserId,
		Role:        role,
	}).ON_CONFLICT(WorkspaceAccess.UserID, WorkspaceAccess.WorkspaceID).DO_UPDATE(
		SET(WorkspaceAccess.Role.SET(String(in.Access.Role))),
	)

	_, err = smt.Exec(lib.DB)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}

	lib.SetWorkspaceRole(utils.AuthAccess{U: userID, W: in.Access.WorkspaceId}, role)

	return &fincal.AddUserToWorkspaceResponse{
		Success: true,
		Error:   "",
	}, nil
}

func (s *WorkspaceServer) RemoveUserFromWorkspace(c context.Context, in *fincal.RemoveUserFromWorkspaceRequest) (*fincal.RemoveUserFromWorkspaceResponse, error) {
	err := protovalidate.Validate(in)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}

	userID, ok := c.Value(utils.UserID{}).(int64)

	if !ok {
		return nil, status.Error(codes.Internal, "Error: Failed to get user id")
	}

	role := lib.GetWorkspaceRole(&utils.AuthAccess{U: userID, W: in.WorkspaceId})

	if in.UserId == userID {
		return nil, status.Errorf(codes.PermissionDenied, "USER_CANNOT_MANAGE_SELF")
	}

	if role != model.Role_Owner {
		return nil, status.Errorf(codes.PermissionDenied, "NOT_OWNER")
	}

	smt := WorkspaceAccess.DELETE().WHERE(
		AND(WorkspaceAccess.UserID.EQ(Int64(in.UserId)), WorkspaceAccess.WorkspaceID.EQ(Int64(in.WorkspaceId))),
	)

	_, err = smt.Exec(lib.DB)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %v", err)
	}

	return &fincal.RemoveUserFromWorkspaceResponse{
		Success: true,
		Error:   "",
	}, nil
}
