// service/service.go
package service

import (
	"errors"

	"github.com/project-app-restaurant/model"
	"github.com/project-app-restaurant/repository"
	"github.com/project-app-restaurant/utils"
)

type Service struct {
	Repo repository.Repository
}

// func (s *Service) RegisterUser(username, password, role string) error {
// 	if role != "Admin" && role != "Koki" && role != "Pelanggan" {
// 		return errors.New("invalid role; must be Admin, Koki, or Pelanggan")
// 	}

// 	user := &model.User{
// 		Username: username,
// 		Password: password,
// 		Role:     role,
// 	}

// 	return s.Repo.RegisterUser(user)
// }

func (s *Service) Login(username, password string) (*model.User, error) {
	user, err := s.Repo.GetUserByUsername(username)
	if err != nil || user.Password != password {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

func (s *Service) AddOrder(user *model.User, order *model.Order) error {
	if user.Role != "Pelanggan" {
		return errors.New("unauthorized role for creating orders")
	}

	if order.DiscountCode != "" {
		discountedPrice, err := utils.ApplyDiscount(order.DiscountCode, order.TotalPrice)
		if err != nil {
			return err
		}
		order.TotalPrice = discountedPrice
	}

	return s.Repo.AddOrder(order)
}

func (s *Service) GetOrders(user *model.User) ([]model.Order, error) {
	if user.Role == "Pelanggan" {
		return nil, errors.New("unauthorized to view orders")
	}
	return s.Repo.GetOrders()
}

func (s *Service) UpdateOrderStatus(user *model.User, orderID int, status string) error {
	if user.Role == "Pelanggan" {
		return errors.New("unauthorized to update orders")
	}
	return s.Repo.UpdateOrderStatus(orderID, status)
}

func (s *Service) DeleteOrder(user *model.User, orderID int) error {
	if user.Role != "Admin" {
		return errors.New("unauthorized to delete orders")
	}
	return s.Repo.DeleteOrder(orderID)
}
