package http

import (
	"fmt"
	"net/http"
	"strconv"

	"fincal-server/.gen/fincal/public/model"

	"github.com/gin-gonic/gin"
)

type Server struct{}

type BadRequestResponse struct {
	Success bool
	Error   string
}

func (s *Server) sendBadRequest(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, BadRequestResponse{
		Success: false,
		Error:   err.Error(),
	})
}

func (s *Server) sendInternalServerError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, BadRequestResponse{
		Success: false,
		Error:   err.Error(),
	})
}

func (s *Server) sendUnauthorized(c *gin.Context, role model.Role) {
	c.JSON(http.StatusUnauthorized, BadRequestResponse{
		Success: false,
		Error:   fmt.Sprintf("Unauthorized. Required role: %s", role),
	})
}

type Auth struct {
	UserId      int64
	WorkspaceId int64
	Role        model.Role
}

func (s *Server) getAuth(c *gin.Context) *Auth {

	us := c.GetHeader("User-Id")
	ws := c.GetHeader("Workspace-Id")

	userId, _ := strconv.ParseInt(us, 10, 64)
	workspaceId, _ := strconv.ParseInt(ws, 10, 64)

	return &Auth{
		UserId:      userId,
		WorkspaceId: workspaceId,
		Role:        model.Role_Owner,
	}
}

func CreateServer() {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	{
		g := v1.Group("/user")

		g.POST()
	}

	r.Run(":7655")
}
