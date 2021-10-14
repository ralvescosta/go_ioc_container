package main

import (
	"ioc/golobby_container/pkg/app"
	"ioc/golobby_container/pkg/domain"
	"ioc/golobby_container/pkg/infra"
	"ioc/golobby_container/pkg/interfaces"
)

func main() {
	container := infra.Newcontainer()

	if err := container.Singleton(
		func() app.IRepository {
			return infra.NewRepository()
		},
	); err != nil {
		panic(err)
	}
	if err := container.Transient(
		func() domain.ISomething {
			return app.NewDoSomething(container)
		},
	); err != nil {
		panic(err)
	}
	container.Singleton(
		func() interfaces.IInterface {
			return interfaces.NewInterface(container)
		},
	)

	var inter interfaces.IInterface
	container.Resolve(&inter)
	inter.Controller(domain.Dto{})
}
