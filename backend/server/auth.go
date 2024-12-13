package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (s *Server) Login(c echo.Context) error {
	var req LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse(err, "Invalid Email or Password"))
	}

	if err := c.Validate(&req); err != nil {

		return c.JSON(http.StatusInternalServerError, ErrorResponse(err, "Failed Validation"))
	}

	resp := LoginResponse{
		AccessToken:  "",
		RefreshToken: "",
	}

	return c.JSON(http.StatusOK, resp)
}

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *Server) SignUp(c echo.Context) error {
	var req LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	}

	return c.JSON(http.StatusOK, LoginResponse{})
}
