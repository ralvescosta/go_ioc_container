package app

import (
	"ioc/golobby_container/pkg/domain"
	"log"
)

type Something struct {
	container IContainer
}

func (pst Something) DoSomething(dto domain.Dto) domain.Entity {
	log.Println("[App::Something::DoSomething]")

	var repository IRepository
	pst.container.Resolve(&repository)

	return repository.Create(dto)
}

func NewDoSomething(container IContainer) domain.ISomething {
	return Something{
		container: container,
	}
}
