package middlewares

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func Auth() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			//.鉴权逻辑写在这里
			c := ctx.(http.Context)
			//.登录接口无需鉴权
			if c.Request().URL.String() == "/ping"{
				return handler(ctx, req)
			}
			//token := c.Request().Header.Get("Authorization")
			//if token == "" {
			//	return &v1.BaseReply{Data: "",Code: 450,Msg: "请先登录"},nil
			//}
			//user,err := autowire.App.Uc.GetUserInfo(token)
			//if err != nil {
			//	return &v1.BaseReply{Data: "",Code: 450,Msg: "请先登录"},nil
			//}
			//userB,_ := json.Marshal(user)
			//c.Request().Header.Set("auth-user",string(userB))
			return handler(ctx, req)
		}
	}
}
