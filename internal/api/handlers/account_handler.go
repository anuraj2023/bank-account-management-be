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


func (h *AccountHandler) GetAccount(c echo.Context) error {
	accountNumber := c.Param("accountNumber")

	account, err := h.repo.GetAccount(c.Request().Context(), accountNumber)
	if err != nil {
		if err == repository.ErrAccountNotFound {
			return echo.NewHTTPError(http.StatusNotFound, "Account not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve account")
	}

	return c.JSON(http.StatusOK, account)
}



func (h *AccountHandler) ListAccounts(c echo.Context) error {
	accounts, err := h.repo.ListAccounts(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve accounts")
	}

	return c.JSON(http.StatusOK, accounts)
}