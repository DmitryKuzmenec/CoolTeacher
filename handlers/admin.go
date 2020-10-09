package handlers

import "github.com/labstack/echo"

// Admin - struct for admin handler
type Admin struct {
}

// NewAdmin - constructor for admin handler
func NewAdmin() *Admin {
	return &Admin{}
}

// Home - handler of root page
func (a *Admin) Home(ctx echo.Context) error {
	return nil
}
