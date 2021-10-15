package main

import (
	"ioc/uber_dig/pkg/app"
	"ioc/uber_dig/pkg/domain"
	"ioc/uber_dig/pkg/infra"
	"ioc/uber_dig/pkg/interfaces"
	"log"

	"go.uber.org/dig"
)

func main() {
	container := dig.New()

	if err := container.Provide(func() app.IRepository {
		return infra.NewRepository()
	}); err != nil {
		log.Fatal(err)
	}
	if err := container.Provide(app.NewDoSomething); err != nil {
		log.Fatal(err)
	}
	if err := container.Provide(interfaces.NewInterface); err != nil {
		log.Fatal(err)
	}

	if err := container.Invoke(func(interfaces interfaces.IInterface) {
		interfaces.Controller(domain.Dto{})
	}); err != nil {
		log.Fatal(err)
	}
}
