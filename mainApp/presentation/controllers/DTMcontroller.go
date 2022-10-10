package controllers

import (
	"encoding/json"
	"github.com/dtm-labs/dtm/client/dtmcli"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

var dtmCoordinatorAddress = "http://localhost:36789/api/dtmsvr"
var ordersServerURL = " http://localhost:8080"
var customersServerURL = "http://localhost:8081"
var mysqlURL = "saga:saga@tcp(mysql:3306)/saga?charset=utf8mb4&parseTime=True&loc=Local"

type DTMController struct {
}

type order struct {
	gorm.Model
	IDTransaction string
	IDCustomer    uint
	Email         string
	Status        string
}

func (dtm *DTMController) createOrder(createOrderRequest *orderRequest) error {

	globalTransactionId := dtmcli.MustGenGid(dtmCoordinatorAddress)
	req, _ := StructToMap(createOrderRequest)

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

func (dtm *DTMController) RegisterOrder(c *gin.Context) {
	registerOrderRequest := struct {
		IdCustomer uint   `json:"idCustomer"`
		Email      string `json:"email"`
	}{}
	transactionId := c.Query("gid")

	err := c.BindJSON(&registerOrderRequest)
	if err != nil {
		c.String(http.StatusInternalServerError, "Помилка сервера")
	}

	//return
	err = getDb().Create(&order{
		IDTransaction: transactionId,
		IDCustomer:    registerOrderRequest.IdCustomer,
		Email:         registerOrderRequest.Email,
		Status:        "created",
	}).Error

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		c.String(http.StatusOK, err.Error())
	}

}

func (dtm *DTMController) RegisterOrderCompensate(c *gin.Context) {
	transactionId := c.Query("gid")

	err := getDb().Model(&order{}).
		Where("id_transaction = ?", transactionId).
		Update("status", "canceled").
		Limit(1).
		Error

	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	}
	c.String(http.StatusOK, err.Error())
}

func StructToMap(obj interface{}) (newMap map[string]interface{}, err error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &newMap)
	return
}

func getDb() *gorm.DB {
	db, err := gorm.Open(mysql.Open(mysqlURL))
	if err != nil {
		panic("failed to connect database")
	}
	_ = db.AutoMigrate(&order{})

	return db
}
