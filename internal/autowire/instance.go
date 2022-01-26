package autowire

import (
	"github.com/google/wire"
	"sweet/internal/biz"
	"sweet/internal/data"
)

var ProviderSet = wire.NewSet(NewInstance)

var App *Instance

type Instance struct {
	Data *data.Data
	Uc *biz.SweetUsecase
}

func NewInstance(data *data.Data,usecase *biz.SweetUsecase) *Instance{
	var app Instance
	app.Uc = usecase
	app.Data = data
	return &app
}
