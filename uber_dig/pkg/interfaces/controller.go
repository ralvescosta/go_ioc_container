package interfaces

import (
	"ioc/uber_dig/pkg/domain"
	"log"
)

type IInterface interface {
	Controller(dto domain.Dto)
}

type Interface struct {
	usecase domain.ISomething
}

func (pst Interface) Controller(dto domain.Dto) {
	log.Println("[Interfaces::Interface::Controller]")
	pst.usecase.DoSomething(dto)
}

func NewInterface(usecase domain.ISomething) IInterface {
	return Interface{
		usecase: usecase,
	}
}
