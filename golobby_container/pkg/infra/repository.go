package infra

import (
	"ioc/golobby_container/pkg/app"
	"ioc/golobby_container/pkg/domain"
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
