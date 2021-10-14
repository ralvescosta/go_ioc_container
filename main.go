package main

import (
	"ioc/pkg/app"
	"ioc/pkg/domain"
	"ioc/pkg/infra"
)

func main() {
	container := infra.Newcontainer()

	container.Singleton(
		func() app.IRepository {
			return infra.NewRepository()
		},
	)
	container.Transient(
		func() domain.ISomething {
			return app.NewDoSomething(container)
		},
	)
}
