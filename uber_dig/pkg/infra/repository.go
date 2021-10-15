package infra

import (
	"ioc/uber_dig/pkg/app"
	"ioc/uber_dig/pkg/domain"
	"log"
)

type Repository struct{}

func (Repository) Create(dto domain.Dto) domain.Entity {
	log.Println("[Infra::Repository::Create]")
	return domain.Entity{}
}

func NewRepository() app.IRepository {
	return Repository{}
}
