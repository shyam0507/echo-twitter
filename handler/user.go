package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/shyam0507/echo-twitter/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	collection = "users"
)

func (h *Handler) Signup(c echo.Context) error {
	u := &model.User{ID: primitive.NewObjectID()}
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	fmt.Println(u)

	// Validate
	if u.Email == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}

	// Save user
	if _, err := h.DB.Collection(collection).InsertOne(context.TODO(), u); err != nil {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}

	fmt.Println("User created: ", u.ID.Hex())

	return c.JSON(http.StatusOK, u)
}
