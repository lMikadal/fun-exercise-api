package postgres

import (
	"time"

	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
)

type Wallet struct {
	ID         int       `postgres:"id"`
	UserID     int       `postgres:"user_id"`
	UserName   string    `postgres:"user_name"`
	WalletName string    `postgres:"wallet_name"`
	WalletType string    `postgres:"wallet_type"`
	Balance    float64   `postgres:"balance"`
	CreatedAt  time.Time `postgres:"created_at"`
}

func (p *Postgres) Wallets() ([]wallet.Wallet, error) {
	rows, err := p.Db.Query("SELECT * FROM user_wallet")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []wallet.Wallet
	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet.Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}
	return wallets, nil
}

func (p *Postgres) WalletsByUserID(id int) ([]wallet.Wallet, error) {
	rows, err := p.Db.Query("SELECT * FROM user_wallet WHERE user_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []wallet.Wallet
	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet.Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}
	return wallets, nil
}

func (p *Postgres) CreateWallet(w wallet.Wallet) error {

	row := p.Db.QueryRow("INSERT INTO user_wallet (user_id, user_name, wallet_name, wallet_type, balance, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, user_id, user_name, wallet_name, wallet_type, balance, created_at",
		w.UserID, w.UserName, w.WalletName, w.WalletType, w.Balance, w.CreatedAt,
	)
	if err := row.Scan(&w.ID, &w.UserID, &w.UserName, &w.WalletName, &w.WalletType, &w.Balance, &w.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (p *Postgres) UpdateWallet(w wallet.Wallet) error {
	_, err := p.Db.Exec("UPDATE user_wallet SET user_id = $1, user_name = $2, wallet_name = $3, wallet_type = $4, balance = $5, created_at = $6 WHERE id = $7",
		w.UserID, w.UserName, w.WalletName, w.WalletType, w.Balance, w.CreatedAt, w.ID,
	)

	return err
}
