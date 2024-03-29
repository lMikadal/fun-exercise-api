package wallet

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	store Storer
}

type Storer interface {
	Wallets() ([]Wallet, error)
	WalletsByUserID(id int) ([]Wallet, error)
	CreateWallet(w Wallet) (Wallet, error)
	UpdateWallet(w Wallet) error
}

func New(db Storer) *Handler {
	return &Handler{store: db}
}

type Err struct {
	Message string `json:"message"`
}

// WalletHandler
// @Summary		Get all wallets
// @Description	Get all wallets
// @Tags			wallet
// @Accept			json
// @Produce		json
// @Success		200	{object}	Wallet
// @Router			/api/v1/wallets [get]
// @Param			wallet_type	query	string	false	"Filter by wallet type"
// @Failure		500	{object}	Err
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

// UserWalletHandler
// @Summary		Get all wallets by user id
// @Description	Get all wallets by user id
// @Tags			wallet
// @Accept			json
// @Produce		json
// @Success		200	{object}	Wallet
// @Router			/api/v1/users/{id}/wallets [get]
// @Param			id	path	int	true	"User ID"
// @Failure		500	{object}	Err
// @Failure		400	{object}	Err
func (h *Handler) UserWalletHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "invalid user id"})
	}

	wallets, err := h.store.WalletsByUserID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, wallets)
}

// CreateWalletHandler
// @Summary		Create wallet
// @Description	Create wallet
// @Tags			wallet
// @Accept			json
// @Produce		json
// @Success		201	{object}	object
// @Router			/api/v1/wallets [post]
// @Param			wallet	body	Wallet	true	"Wallet object"
// @Failure		500	{object}	Err
// @Failure		400	{object}	Err
func (h *Handler) CreateWalletHandler(c echo.Context) error {
	w := Wallet{}
	if err := c.Bind(&w); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	wallet, err := h.store.CreateWallet(w)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, wallet)
}

func (h *Handler) UpdateWalletHandler(c echo.Context) error {
	w := Wallet{}
	if err := c.Bind(&w); err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "invalid wallet id"})
	}

	w.ID = id
	err = h.store.UpdateWallet(w)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, w)
}
