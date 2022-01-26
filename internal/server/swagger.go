package server

import (
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	"sweet/internal/conf"
)

func NewSwaggerServer(c *conf.Server,srv *http.Server){
	if c.Swagger.GetPower() == false {
		return
	}
	//.@volta 使用/q/swagger-ui/可访问
	openAPIhandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIhandler)
}
