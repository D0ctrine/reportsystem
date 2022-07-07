package dblayer

import (
	"report/models"
)

type DBLayer interface {
	GetAllProducts() ([]models.Product, error)
	GetProduct(uint) (models.Product, error)
}
