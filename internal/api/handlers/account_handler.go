package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/anuraj2023/bank-account-management-be/internal/models"
	"github.com/anuraj2023/bank-account-management-be/internal/repository"
)

type AccountHandler struct {
	repo repository.AccountRepository
}

func NewAccountHandler(repo repository.AccountRepository) *AccountHandler {
	return &AccountHandler{repo: repo}
}

// @Summary Create a new account
// @Description Create a new bank account
// @Tags accounts
// @Accept json
// @Produce json
// @Param account body models.Account true "Account details"
// @Success 201 {object} models.Account
// @Failure 400 {object} echo.HTTPError
// @Failure 500 {object} echo.HTTPError
// @Router /accounts [post]
func (h *AccountHandler) CreateAccount(c echo.Context) error {
	var account models.Account
	if err := c.Bind(&account); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request payload")
	}

	if err := account.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	createdAccount, err := h.repo.CreateAccount(c.Request().Context(), &account)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create account")
	}

	return c.JSON(http.StatusCreated, createdAccount)
}


// @Summary List all accounts
// @Description Retrieve a list of all bank accounts
// @Tags accounts
// @Produce json
// @Success 200 {array} models.Account
// @Failure 500 {object} echo.HTTPError
// @Router /accounts [get]
func (h *AccountHandler) GetAllAccounts(c echo.Context) error {
	accounts, err := h.repo.GetAllAccounts(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve accounts")
	}

	return c.JSON(http.StatusOK, accounts)
}