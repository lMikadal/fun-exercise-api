package wallet

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	store Storer
}

type Storer interface {
	Wallets() ([]Wallet, error)
}

func New(db Storer) *Handler {
	return &Handler{store: db}
}

type Err struct {
	Message string `json:"message"`
}

// WalletHandler
//
//	@Summary		Get all wallets
//	@Description	Get all wallets
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Wallet
//	@Router			/api/v1/wallets [get]
//	@Param			wallet_type	query	string	false	"Filter by wallet type"
//	@Failure		500	{object}	Err
func (h *Handler) WalletHandler(c echo.Context) error {
	wallets, err := h.store.Wallets()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	w_type := c.QueryParam("wallet_type")
	if w_type == "" {
		return c.JSON(http.StatusOK, wallets)
	}

	var filteredWallets []Wallet
	for _, w := range wallets {
		if w.WalletType == w_type {
			filteredWallets = append(filteredWallets, w)
		}
	}

	return c.JSON(http.StatusOK, filteredWallets)
}

func (h *Handler) UserWalletHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "user wallet handler")
}
