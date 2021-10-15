package app

import (
	"ioc/uber_dig/pkg/domain"
	"log"
)

type Something struct {
	repository IRepository
}

func (pst Something) DoSomething(dto domain.Dto) domain.Entity {
	log.Println("[App::Something::DoSomething]")
	return pst.repository.Create(dto)
}

func NewDoSomething(repository IRepository) domain.ISomething {
	return Something{
		repository: repository,
	}
}
