package app

import "ioc/pkg/domain"

type IRepository interface {
	Create(dto domain.Dto) domain.Entity
}

type IContainer interface {
	Resolve(instance interface{}) interface{}
	NamedResolve(instance interface{}, name string) interface{}
	Singleton(resolver interface{}) error
	NamedSingleton(name string, resolver interface{}) error
	Transient(resolver interface{}) error
	NamedTransient(name string, resolver interface{}) error
}
