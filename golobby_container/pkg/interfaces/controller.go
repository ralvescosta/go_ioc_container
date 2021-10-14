package interfaces

import (
	"ioc/golobby_container/pkg/app"
	"ioc/golobby_container/pkg/domain"
	"log"
)

type IInterface interface {
	Controller(dto domain.Dto)
}

type Interface struct {
	container app.IContainer
}

func (pst Interface) Controller(dto domain.Dto) {
	log.Println("[Interfaces::Interface::Controller]")

	var usecase domain.ISomething
	pst.container.Resolve(&usecase)

	usecase.DoSomething(dto)
}

func NewInterface(container app.IContainer) Interface {
	return Interface{
		container: container,
	}
}
