package interfaces

import (
	"ioc/pkg/app"
	"ioc/pkg/domain"
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
