package lib

import (
	"fincal-server/.gen/fincal/public/model"
	. "fincal-server/.gen/fincal/public/table"
	"fincal-server/utils"

	. "github.com/go-jet/jet/v2/postgres"

	lru "github.com/hashicorp/golang-lru/v2"
)

var (
	AuthCache *lru.Cache[utils.AuthAccess, model.Role]
)

func init() {
	l, err := lru.New[utils.AuthAccess, model.Role](1024)

	if err != nil {
		Sugar.Fatal(err)
	}

	AuthCache = l
}

func GetWorkspaceRole(in *utils.AuthAccess) model.Role {
	role, ok := AuthCache.Get(*in)

	if ok {
		return role
	}

	stmt := WorkspaceAccess.SELECT(WorkspaceAccess.Role).WHERE(
		AND(WorkspaceAccess.UserID.EQ(Int64(in.U)), WorkspaceAccess.WorkspaceID.EQ(Int64(in.W))),
	)

	var dest []struct {
		model.WorkspaceAccess
	}

	err := stmt.Query(DB, &dest)

	if err != nil {
		return model.Role_None
	}

	if len(dest) == 0 {
		role = model.Role_None
	}

	AuthCache.Add(*in, role)

	return model.Role_Read
}

func SetWorkspaceRole(in utils.AuthAccess, role model.Role) bool {
	return AuthCache.Add(in, role)
}
