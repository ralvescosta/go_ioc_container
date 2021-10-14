package main

import (
	"ioc/pkg/app"
	"ioc/pkg/domain"
	"ioc/pkg/infra"
	"ioc/pkg/interfaces"
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
