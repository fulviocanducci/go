package database

import (
	"database/sql"

	"github.com/fulviocanducci/go/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB

	// Save(order *Order) error
	// GetTotal() (int, error)
}

func (r *OrderRepository) Save(o *entity.Order) error {
	_, err := r.Db.Exec("INSERT INTO orders(id, price, tax, final_price) VALUES(?,?,?,?)", o.Id, o.Price, o.Tax, o.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotalTransactions() (int, error) {
	var total int
	err := r.Db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
