package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	Type      string
	Email     string
	Password  string
	GCode     string
	AppleCode string
}

type CreateUserOutput struct {
	Success bool
	Error   string
	Token   string
}

func (s *Server) CreateUser(c *gin.Context) {
	var input CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, CreateUserOutput{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	// TODO: Create Token

	c.JSON(http.StatusOK, CreateUserOutput{
		Success: true,
		Token:   "token",
	})
}
