package services

import (
	"encoding/json"
	"errors"
	"github.com/dtm-labs/dtm/client/dtmcli"
	"github.com/vsnnkv/btcApplicationGo/infrastructure/repository"
	"net/mail"
)

type SubscriptionServiceInterface interface {
	SaveEmail(email string) error
	CreateCustomer(email string) error
}

type SubscriptionService struct {
	emailService EmailService
	customerDB   repository.DtmDbInterface
}

func NewSubscriptionService(f EmailService, db repository.DtmDbInterface) *SubscriptionService {
	return &SubscriptionService{emailService: f, customerDB: db}
}

type orderRequest struct {
	IdCustomer uint   `json:"idCustomer"`
	Currency   string `json:"currency"`
	Amount     uint   `json:"amount"`
}

var dtmCoordinatorAddress = "http://localhost:36789/api/dtmsvr"
var ordersServerURL = " http://localhost:8080"
var customersServerURL = "http://localhost:8081"

func (s *SubscriptionService) SaveEmail(email string) error {
	if !isEmailValid(email) {
		return errors.New("email not valid")
	}
	exist, err := s.emailService.repository.IsExists(email)

	if err != nil {
		return errors.New("Помилка сервера")
	}

	if exist {
		return errors.New("Email вже було додано")
	}

	err = s.emailService.repository.SaveEmailToFile(email)
	if err != nil {
		return errors.New("Помилка збереження файла")
	}
	return nil
}

func (s *SubscriptionService) CreateCustomer(email string) error {

	err, customerId := s.customerDB.SaveCustomerToDb(email)

	if err != nil {
		return err
	}
	customer := orderRequest{
		IdCustomer: customerId,
		Amount:     10,
		Currency:   "BTC",
	}

	globalTransactionId := dtmcli.MustGenGid(dtmCoordinatorAddress)
	req, _ := structToMap(customer)

	_ = dtmcli.
		NewSaga(dtmCoordinatorAddress, globalTransactionId).
		Add(ordersServerURL+"/register-order", ordersServerURL+"/register-order-compensate", req).
		Add(customersServerURL+"/withdraw-money", customersServerURL+"/withdraw-money-compensate", req).
		Submit()

	_ = struct {
		Gid string `json:"gid"`
	}{Gid: globalTransactionId}

	return nil
}

func isEmailValid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func structToMap(obj interface{}) (newMap map[string]interface{}, err error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &newMap)
	return
}
