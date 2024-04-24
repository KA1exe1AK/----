package main

type Controller struct {
	service *Service
	// teacherService
	// student service
}

func NewController(service *Service) *Controller {
	return &Controller{
		service: service,
	}
}
