package infra

import (
	"ioc/pkg/app"
	"ioc/pkg/domain"
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
