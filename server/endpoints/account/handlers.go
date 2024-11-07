package account

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var dao = NewInMemory()

func getAccountByIDHandler(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, "Invalid address ID")
	}

	acc, err := dao.GetAccountByID(id)

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, acc)
}

type createAccountRequestBody struct {
	Email string `json:"email"`
}

func createAccountHandler(c echo.Context) error {
	body := new(createAccountRequestBody)
	if err := c.Bind(&body); err != nil {
		return c.String(http.StatusBadRequest, "Bad request")
	}

	acc, err := dao.CreateAccount(body.Email)
	if err != nil {
		return c.String(http.StatusInternalServerError, "An error occurred while creating account")
	}

	return c.JSON(http.StatusOK, acc)
}

func RegisterHandlers(e *echo.Echo) {
	e.GET("/account/:id", getAccountByIDHandler)
	e.POST("/account", createAccountHandler)
}
