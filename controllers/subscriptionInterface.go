package controllers

type SubscriptionService interface {
	SaveEmail(email string) (int, string)
}

type SubscriptionController struct {
	subscriptionService SubscriptionService
}

func NewSubscriptionController(s SubscriptionService) *SubscriptionController {
	return &SubscriptionController{subscriptionService: s}
}
