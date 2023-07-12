package api

import (
	"github.com/cpglsn/hotel-reservation/types"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.userStore
}

func NewUserHandler(userStore db.userStore) *UserHandler {
	return &UserUserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.userStore.GetUserByID(id)
	if err != nil {
		return err
	}
	c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName:  "Watercooler",
	}
	return c.JSON(u)
}
