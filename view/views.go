package view

import (
	"encoding/json"
	"fmt"

	"github.com/project-app-restaurant/model"
)

func DisplayOrders(orders []model.Order) {
	data, err := json.MarshalIndent(orders, "", "  ")
	if err != nil {
		fmt.Println("Error formatting orders:", err)
		return
	}
	fmt.Println(string(data))
}
