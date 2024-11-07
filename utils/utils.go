package utils

import "errors"

func ApplyDiscount(code string, price float64) (float64, error) {
	discounts := map[string]float64{
		"DISCOUNT10": 0.10,
		"DISCOUNT20": 0.20,
	}
	discount, exists := discounts[code]
	if !exists {
		return price, errors.New("invalid discount code")
	}
	return price * (1 - discount), nil
}
