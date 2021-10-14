package app

import "ioc/pkg/domain"

type IRepository interface {
	Create(dto domain.Dto) domain.Entity
}
