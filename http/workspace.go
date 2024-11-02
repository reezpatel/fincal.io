package http

import (
	"net/http"

	"fincal-server/.gen/fincal/public/model"
	. "fincal-server/.gen/fincal/public/table"
	"fincal-server/lib"

	"github.com/gin-gonic/gin"
	"github.com/go-jet/jet/v2/postgres"
)

type CreateWorkspaceInput struct {
	Name         string
	CurrencyCode string
	Icon         string
}

type CreateWorkspaceOutput struct {
	WorkspaceId int64
}

func (s *Server) CreateWorkspace(c *gin.Context) {
	auth := s.getAuth(c)

	var input CreateWorkspaceInput

	if err := c.ShouldBindJSON(&input); err != nil {
		s.sendBadRequest(c, err)
		return
	}

	smt := Workspace.INSERT(Workspace.MutableColumns).MODEL(model.Workspace{
		Name:         input.Name,
		CurrencyCode: input.CurrencyCode,
		Icon:         input.Icon,
	}).RETURNING(Workspace.ID)

	var dest struct {
		ID int64
	}

	err := smt.Query(lib.DB, &dest)

	if err != nil {
		s.sendInternalServerError(c, err)
		return
	}

	smt2 := WorkspaceAccess.INSERT(WorkspaceAccess.MutableColumns).MODEL(model.WorkspaceAccess{
		WorkspaceID: dest.ID,
		UserID:      auth.UserId,
		Role:        model.Role_Owner,
	})

	_, err = smt2.Exec(lib.DB)

	if err != nil {
		s.sendInternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, CreateWorkspaceOutput{
		WorkspaceId: dest.ID,
	})
}

func (s *Server) UpdateWorkspace(c *gin.Context) {
	auth := s.getAuth(c)

	if auth.Role != model.Role_Owner {
		s.sendUnauthorized(c, model.Role_Owner)
		return
	}

	var input CreateWorkspaceInput

	if err := c.ShouldBindJSON(&input); err != nil {
		s.sendBadRequest(c, err)
		return
	}

	smt := Workspace.UPDATE(Workspace.MutableColumns).MODEL(model.Workspace{
		Name:         input.Name,
		CurrencyCode: input.CurrencyCode,
		Icon:         input.Icon,
	}).WHERE(Workspace.ID.EQ(postgres.Int64(auth.WorkspaceId)))

	_, err := smt.Exec(lib.DB)

	if err != nil {
		s.sendInternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (s *Server) GetWorkspaces(c *gin.Context) {
	auth := s.getAuth(c)

	stmt := postgres.SELECT(Workspace.AllColumns).FROM(
		WorkspaceAccess.INNER_JOIN(Workspace, Workspace.ID.EQ(WorkspaceAccess.WorkspaceID))).WHERE(
		Workspace.ID.EQ(postgres.Int64(auth.WorkspaceId)),
	)

	var dest []model.Workspace

	err := stmt.Query(lib.DB, &dest)

	if err != nil {
		s.sendInternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, dest)
}

func (s *Server) RemoveWorkspace(c *gin.Context) {
	auth := s.getAuth(c)

	if auth.Role != model.Role_Owner {
		s.sendUnauthorized(c, model.Role_Owner)
		return
	}

	smt := Workspace.DELETE().WHERE(Workspace.ID.EQ(postgres.Int64(auth.WorkspaceId)))

	_, err := smt.Exec(lib.DB)

	if err != nil {
		s.sendInternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}
