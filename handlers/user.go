package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/DmitryKuzmenec/CoolTeacher/logger"
	"github.com/DmitryKuzmenec/CoolTeacher/model"
	"github.com/DmitryKuzmenec/CoolTeacher/services"
	"github.com/labstack/echo"
)

// UserHandler - handler of user
type UserHandler struct {
	user *services.UserService
	log  *logger.Logger
}

// NewUserHandler - constructor of UserHandler
func NewUserHandler(user *services.UserService, log *logger.Logger) *UserHandler {
	return &UserHandler{
		user: user,
		log:  log,
	}
}

// Create - create user
func (h *UserHandler) Create(ctx echo.Context) error {
	var user model.User
	if err := ctx.Bind(&user); err != nil {
		h.log.Error(err)
		return nil
	}
	if user.Fname == "" {
		h.log.Error(errors.New("user first name is empty!"))
		return ctx.String(http.StatusBadRequest, "User first name is empty!")
	}
	if user.Lname == "" {
		return ctx.String(http.StatusBadRequest, "User last name is empty!")
	}
	if user.Gender == "" {
		return ctx.String(http.StatusBadRequest, "User gender name is empty!")
	}
	id, err := h.user.Create(&user)
	if err != nil {
		return err
	}
	return ctx.String(http.StatusOK, fmt.Sprintf("UserID: %d", id))
}

// List - list of all users
func (h *UserHandler) List(ctx echo.Context) error {
	users, err := h.user.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &users)
}
