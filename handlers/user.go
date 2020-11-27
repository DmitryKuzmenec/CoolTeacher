package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"time"

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
	var userReq model.UserData
	if err := ctx.Bind(&userReq); err != nil {
		h.log.Error(err)
		return nil
	}
	if userReq.Fname == "" {
		h.log.Error(errors.New("user first name is empty!"))
		return ctx.String(http.StatusBadRequest, "User first name is empty!")
	}
	if userReq.Lname == "" {
		return ctx.String(http.StatusBadRequest, "User last name is empty!")
	}
	if userReq.Gender == "" {
		return ctx.String(http.StatusBadRequest, "User gender name is empty!")
	}
	birthDate := time.Time(userReq.BirthDate)
	user := model.User{
		Fname:       userReq.Fname,
		Lname:       userReq.Lname,
		Gender:      userReq.Gender,
		BirthDate:   &birthDate,
		Phone:       userReq.Phone,
		Email:       userReq.Email,
		Description: userReq.Description,
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
