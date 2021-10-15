package app

import "ioc/uber_dig/pkg/domain"

type IRepository interface {
	Create(dto domain.Dto) domain.Entity
}
