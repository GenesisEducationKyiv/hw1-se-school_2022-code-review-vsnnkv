package services

import "github.com/vsnnkv/btcApplicationGo/infrastructure/repository"

type DTMServiceInterface interface {
	ChangeCustomerStatus(transactionId string, customerId uint, currency string, amount uint) error
	ChangeCustomerStatusCompensation(transactionId string) error
}

type DTMService struct {
	db repository.DtmDbInterface
}

func NewDtmService(d repository.DtmDbInterface) *DTMService {
	return &DTMService{db: d}
}

func (service *DTMService) ChangeCustomerStatus(transactionId string, customerId uint,
	currency string, amount uint) error {

	return service.db.CreateTransaction(transactionId, customerId, currency, amount)
}

func (service *DTMService) ChangeCustomerStatusCompensation(transactionId string) error {

	return service.db.TransactionCompensation(transactionId)
}