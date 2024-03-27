package wallet

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
)

type StubWallet struct {
	wallet []Wallet
	err    error
}

func (s StubWallet) Wallets() ([]Wallet, error) {
	return s.wallet, s.err
}

func TestWallet(t *testing.T) {
	t.Run("given unable to get wallets should return 500 and error message", func(t *testing.T) {
		c, rec := request(http.MethodGet, "/api/v1/wallets", nil)

		stubError := StubWallet{err: echo.ErrInternalServerError}
		w := New(stubError)

		w.WalletHandler(c)

		if rec.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %d but got %d", http.StatusInternalServerError, rec.Code)
		}
	})

	t.Run("given user able to getting wallet should return list of wallets", func(t *testing.T) {
		c, rec := request(http.MethodGet, "/api/v1/wallets", nil)

		timeNow, err := time.Parse(time.RFC3339Nano, "2024-03-25T14:19:00.729237Z")
		if err != nil {
			t.Errorf("unable to parse time")
		}
		stubWallets := StubWallet{
			wallet: []Wallet{
				{ID: 1, UserID: 1, UserName: "John Doe", WalletName: "John's Wallet", WalletType: "Create Card", Balance: 100.00, CreatedAt: timeNow},
				{ID: 2, UserID: 2, UserName: "Jane Doe", WalletName: "Jane's Wallet", WalletType: "Create Card", Balance: 200.00, CreatedAt: timeNow},
			},
			err: nil,
		}

		p := New(stubWallets)
		p.WalletHandler(c)

		want := []Wallet{
			{ID: 1, UserID: 1, UserName: "John Doe", WalletName: "John's Wallet", WalletType: "Create Card", Balance: 100.00, CreatedAt: timeNow},
			{ID: 2, UserID: 2, UserName: "Jane Doe", WalletName: "Jane's Wallet", WalletType: "Create Card", Balance: 200.00, CreatedAt: timeNow},
		}
		gotJson := rec.Body.Bytes()
		var got []Wallet
		if err := json.Unmarshal(gotJson, &got); err != nil {
			t.Errorf("unable to unmarshal response body")
		}
		if rec.Code != http.StatusOK {
			t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v but got %v", want, got)
		}
	})

	t.Run("given user able to getting wallet with filter type Savings should return list of wallets have type Savings", func(t *testing.T) {
		c, rec := request(http.MethodGet, "/api/v1/wallets?wallet_type=Savings", nil)

		timeNow, err := time.Parse(time.RFC3339Nano, "2024-03-25T14:19:00.729237Z")
		if err != nil {
			t.Errorf("unable to parse time")
		}
		stubWallets := StubWallet{
			wallet: []Wallet{
				{ID: 1, UserID: 1, UserName: "John Doe", WalletName: "John's Wallet", WalletType: "Savings", Balance: 100.00, CreatedAt: timeNow},
				{ID: 2, UserID: 2, UserName: "Jane Doe", WalletName: "Jane's Wallet", WalletType: "Create Card", Balance: 200.00, CreatedAt: timeNow},
			},
			err: nil,
		}

		p := New(stubWallets)
		p.WalletHandler(c)

		want := []Wallet{
			{ID: 1, UserID: 1, UserName: "John Doe", WalletName: "John's Wallet", WalletType: "Savings", Balance: 100.00, CreatedAt: timeNow},
		}
		gotJson := rec.Body.Bytes()
		var got []Wallet
		if err := json.Unmarshal(gotJson, &got); err != nil {
			t.Errorf("unable to unmarshal response body")
		}
		if rec.Code != http.StatusOK {
			t.Errorf("expected status code %d but got %d", http.StatusOK, rec.Code)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v but got %v", want, got)
		}
	})
}

func request(method, path string, body io.Reader) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, path, body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec
}
