// +build memory

package models

import (
	"github.com/alimy/gin-music/models/core"
)

func init() {
	core.RegisterRepository(MemRepository)
}

var (
	MemRepository = &core.DefaultRepository{OnRepository: &memOnRepository{
		template: &memTemplate{},
		Albums:   make(map[string]*Album),
	}}
)

type memOnRepository struct {
	template core.Template
	Albums   map[string]*Album
}

func (r *memOnRepository) OnCreate(model core.RecyclableCrud) error {
	return r.template.Create(model)
}

func (r *memOnRepository) OnRetrieve(model core.RecyclableCrud) error {
	return r.template.Retrieve(model)
}

func (r *memOnRepository) OnUpdate(model core.RecyclableCrud) error {
	return r.template.Update(model)
}

func (r *memOnRepository) OnDelete(model core.RecyclableCrud) error {
	return r.template.Delete(model)
}
