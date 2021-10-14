package infra

import (
	"ioc/golobby_container/pkg/app"

	containerPkg "github.com/golobby/container/v3/pkg/container"
)

type container struct {
	ctx containerPkg.Container
}

func (pst *container) Resolve(instance interface{}) interface{} {
	return pst.ctx.Resolve(instance)
}
func (pst *container) NamedResolve(instance interface{}, name string) interface{} {
	return pst.ctx.NamedResolve(instance, name)
}
func (pst *container) Singleton(resolver interface{}) error {
	return pst.ctx.Singleton(resolver)
}
func (pst *container) NamedSingleton(name string, resolver interface{}) error {
	return pst.ctx.NamedSingleton(name, resolver)
}
func (pst *container) Transient(resolver interface{}) error {
	return pst.ctx.Transient(resolver)
}
func (pst *container) NamedTransient(name string, resolver interface{}) error {
	return pst.ctx.NamedTransient(name, resolver)
}
func Newcontainer() app.IContainer {
	return &container{
		ctx: containerPkg.New(),
	}
}
