package main

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/project-app-restaurant/database"
	"github.com/project-app-restaurant/model"
	"github.com/project-app-restaurant/repository"
	"github.com/project-app-restaurant/service"
	"github.com/project-app-restaurant/view"
)

var currentUser *model.User

func main() {
	db := database.InitDB()
	defer db.Close()

	repo := &repository.PostgresRepository{DB: db}
	svc := &service.Service{Repo: repo}

	for {
		if currentUser == nil {
			loginUser(svc)
		} else {
			switch currentUser.Role {
			case "Admin":
				adminMenu(svc)
			case "Koki":
				kokiMenu(svc)
			case "Pelanggan":
				pelangganMenu(svc)
			}
		}
	}
}

func loginUser(svc *service.Service) {
	var username, password string
	fmt.Print("Username: ")
	fmt.Scan(&username)
	fmt.Print("Password: ")
	fmt.Scan(&password)

	user, err := svc.Login(username, password)
	if err != nil {
		fmt.Println("Login failed:", err)
		return
	}

	currentUser = user
	fmt.Printf("Login successful! Welcome, %s (%s)\n", currentUser.Username, currentUser.Role)
}

func adminMenu(svc *service.Service) {
	fmt.Println("\n-- Admin Menu --")
	fmt.Println("1. Add new order")
	fmt.Println("2. View all orders")
	fmt.Println("3. Update order status")
	fmt.Println("4. Delete order")
	fmt.Println("5. Logout")
	fmt.Print("Choose an option: ")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		addOrder(svc)
	case 2:
		viewOrders(svc)
	case 3:
		updateOrderStatus(svc)
	case 4:
		deleteOrder(svc)
	case 5:
		logout()
	default:
		fmt.Println("Invalid option")
	}
}

func kokiMenu(svc *service.Service) {
	fmt.Println("\n-- Koki Menu --")
	fmt.Println("1. View all orders")
	fmt.Println("2. Update order status")
	fmt.Println("3. Logout")
	fmt.Print("Choose an option: ")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		viewOrders(svc)
	case 2:
		updateOrderStatus(svc)
	case 3:
		logout()
	default:
		fmt.Println("Invalid option")
	}
}

func pelangganMenu(svc *service.Service) {
	fmt.Println("\n-- Pelanggan Menu --")
	fmt.Println("1. Add new order")
	fmt.Println("2. View your orders")
	fmt.Println("3. Logout")
	fmt.Print("Choose an option: ")
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		addOrder(svc)
	case 2:
		viewOrders(svc)
	case 3:
		logout()
	default:
		fmt.Println("Invalid option")
	}
}

func addOrder(svc *service.Service) {
	var totalPrice float64
	var discountCode string

	fmt.Print("Enter total price: ")
	fmt.Scan(&totalPrice)
	fmt.Print("Enter discount code (if any): ")
	fmt.Scan(&discountCode)

	order := &model.Order{
		UserID:       currentUser.ID,
		OrderStatus:  "sedang diproses",
		TotalPrice:   totalPrice,
		DiscountCode: discountCode,
	}

	err := svc.AddOrder(currentUser, order)
	if err != nil {
		fmt.Println("Error adding order:", err)
		return
	}

	fmt.Println("Order added successfully!")
}

func viewOrders(svc *service.Service) {
	orders, err := svc.GetOrders(currentUser)
	if err != nil {
		fmt.Println("Error fetching orders:", err)
		return
	}
	view.DisplayOrders(orders)
}

func updateOrderStatus(svc *service.Service) {
	var orderID int
	var newStatus string

	fmt.Print("Enter order ID to update: ")
	fmt.Scan(&orderID)
	fmt.Print("Enter new status (sedang diproses / selesai / dibatalkan): ")
	fmt.Scan(&newStatus)

	err := svc.UpdateOrderStatus(currentUser, orderID, newStatus)
	if err != nil {
		fmt.Println("Error updating order status:", err)
		return
	}

	fmt.Println("Order status updated successfully!")
}

func deleteOrder(svc *service.Service) {
	if currentUser.Role != "Admin" {
		fmt.Println("Unauthorized access")
		return
	}

	var orderID int
	fmt.Print("Enter order ID to delete: ")
	fmt.Scan(&orderID)

	err := svc.DeleteOrder(currentUser, orderID)
	if err != nil {
		fmt.Println("Error deleting order:", err)
		return
	}

	fmt.Println("Order deleted successfully!")
}

func logout() {
	currentUser = nil
	fmt.Println("Logged out successfully!")
}
