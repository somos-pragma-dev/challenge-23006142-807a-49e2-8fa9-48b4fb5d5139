package utils

import (
	"errors"
	"src/services"
)

func ValidateProductDTO(productDTO services.CreateProductDTO) error {
	if productDTO.Name == "" {
		return errors.New("name is required")
	}
	if productDTO.Price < 0 {
		return errors.New("price must be positive")
	}
	return nil
}