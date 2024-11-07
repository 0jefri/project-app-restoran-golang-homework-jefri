package repository

import (
	"database/sql"
	"log"

	"github.com/project-app-restaurant/model"
)

type Repository interface {
	// RegisterUser(user *model.User) error
	GetUserByUsername(username string) (*model.User, error)
	AddOrder(order *model.Order) error
	GetOrders() ([]model.Order, error)
	UpdateOrderStatus(orderID int, status string) error
	DeleteOrder(orderID int) error
}

type PostgresRepository struct {
	DB *sql.DB
}

// func (r *PostgresRepository) RegisterUser(user *model.User) error {
// 	query := `INSERT INTO users (username, password, role) VALUES ($1, $2, $3)`
// 	_, err := r.DB.Exec(query, user.Username, user.Password, user.Role)
// 	if err != nil {
// 		return errors.New("failed to register user: " + err.Error())
// 	}
// 	return nil
// }

func (repo *PostgresRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	query := `SELECT id, username, role, password FROM users WHERE username=$1`
	err := repo.DB.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Role, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *PostgresRepository) AddOrder(order *model.Order) error {
	query := `INSERT INTO orders (user_id, order_status, total_price, discount_code) VALUES ($1, $2, $3, $4) RETURNING id`
	err := repo.DB.QueryRow(query, order.UserID, order.OrderStatus, order.TotalPrice, order.DiscountCode).Scan(&order.ID)
	return err
}

func (repo *PostgresRepository) GetOrders() ([]model.Order, error) {
	rows, err := repo.DB.Query(`SELECT id, user_id, order_status, total_price, discount_code, rating, created_at FROM orders`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []model.Order
	for rows.Next() {
		var order model.Order
		if err := rows.Scan(&order.ID, &order.UserID, &order.OrderStatus, &order.TotalPrice, &order.DiscountCode, &order.Rating, &order.CreatedAt); err != nil {
			log.Println("Error scanning order:", err)
			continue
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (repo *PostgresRepository) UpdateOrderStatus(orderID int, status string) error {
	query := `UPDATE orders SET order_status=$1 WHERE id=$2`
	_, err := repo.DB.Exec(query, status, orderID)
	return err
}

func (repo *PostgresRepository) DeleteOrder(orderID int) error {
	query := `DELETE FROM orders WHERE id=$1`
	_, err := repo.DB.Exec(query, orderID)
	return err
}
