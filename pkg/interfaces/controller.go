package interfaces

import (
	"ioc/pkg/domain"
	"log"
)

type Interface struct {
	Usecase domain.ISomething
}

func (pst Interface) Controller() {
	log.Println("[Interfaces::Interface::Controller]")
	pst.Usecase.DoSomething(domain.Dto{})
}

func NewInterface() Interface {
	return Interface{}
}
