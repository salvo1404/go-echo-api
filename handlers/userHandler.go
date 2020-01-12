package userHandler

import (
	"net/http"

	"github.com/labstack/echo"
	user "github.com/salvo1404/go-echo-api/models"
)

type (
	resultLists struct {
		Users []user.User `json:"users"`
	}

	handler struct {
		UserModel user.UserModelImpl
	}
)

func NewHandler(u user.UserModelImpl) *handler {
	return &handler{u}
}

func (h *handler) GetIndex(c echo.Context) error {
	lists := h.UserModel.FindAll()

	return c.JSON(http.StatusOK, lists)
}

func (h *handler) GetDetail(c echo.Context) error {
	id := c.Param("id")
	u := h.UserModel.FindByID(id)

	if (user.User{}) == u {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, u)
}

func (h *handler) Save(c echo.Context) error {
	name := c.FormValue("name")

	if len(name) < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing param Name")
	}

	u := h.UserModel.Store(name)
	return c.JSON(http.StatusOK, u)
}

func (h *handler) Delete(c echo.Context) error {
	id := c.Param("id")
	u := h.UserModel.Delete(id)

	if (user.User{}) == u {
		return echo.NewHTTPError(http.StatusNotFound, "User ID does not exist")
	}

	var response struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
	}
	response.Message = "User Deleted"
	response.Code = http.StatusOK

	return c.JSON(response.Code, response)
}
