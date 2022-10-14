package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DtmDbInterface interface {
	SaveCustomerToDb(email string) (error, uint)
	CreateTransaction(transactionId string, customerId uint, currency string, amount uint) error
	TransactionCompensation(transactionId string) error
}

type DtmDb struct {
}

type customer struct {
	gorm.Model
	Email   string
	Balance uint
	Version uint
}

type order struct {
	gorm.Model
	IDTransaction string
	IDCustomer    uint
	Currency      string
	Amount        uint
	Status        string
}

var mysqlURL = "saga:saga@tcp(localhost:8092)/saga?charset=utf8mb4&parseTime=True&loc=Local"

func (*DtmDb) SaveCustomerToDb(email string) (error, uint) {
	customer := customer{Email: email, Balance: 1, Version: 0}

	err := getDb().Save(&customer).Error

	return err, customer.Model.ID
}

func (*DtmDb) CreateTransaction(transactionId string, customerId uint, currency string, amount uint) error {

	return getDbOrder().Create(&order{
		IDTransaction: transactionId,
		IDCustomer:    customerId,
		Currency:      currency,
		Amount:        amount,
		Status:        "created",
	}).Error
}

func (*DtmDb) TransactionCompensation(transactionId string) error {
	return getDbOrder().Model(&order{}).
		Where("id_transaction = ?", transactionId).
		Update("status", "canceled").
		Limit(1).
		Error
}

func getDb() *gorm.DB {
	db, err := gorm.Open(mysql.Open(mysqlURL))
	if err != nil {
		panic("failed to connect database")
	}
	_ = db.AutoMigrate(&customer{})

	return db
}

func getDbOrder() *gorm.DB {
	db, err := gorm.Open(mysql.Open(mysqlURL))
	if err != nil {
		panic("failed to connect database")
	}
	_ = db.AutoMigrate(&order{})

	return db
}
