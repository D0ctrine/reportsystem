package dblayer

import (
	"report/models"

	"github.com/cengsin/oracle"
	"gorm.io/gorm"
)

//get products
//get promos
//post user sign in
//get user orders
//post user sign out
//post purchase charge

type DBORM struct {
	*gorm.DB
}

func NewORM(dbname, con string) (*DBORM, error) {
	//    계정/비밀번호@호스트/명 coms/coms@hmcos01_vip:1521/HMOCS
	db, err := gorm.Open(oracle.Open(dbname+"@"+con), &gorm.Config{})
	return &DBORM{
		DB: db,
	}, err
}

func (db *DBORM) GetAllProducts() (products []models.Product, err error) {
	return products, db.Find(&products).Error
}

func (db *DBORM) GetProduct(uint) (models.Product, error) {
	return models.Product{}, nil
}
