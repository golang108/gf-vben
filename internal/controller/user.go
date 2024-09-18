package controller

import (
	"Gf-Vben/api/v1/user"
	"Gf-Vben/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/jinmao88/gf-utility/response"
)

var (
	User = cUser{}
)

type cUser struct {
}

func (cUser) Register(ctx context.Context, req *user.RegisterReq) (res *response.JsonRes, err error) {
	res = new(response.JsonRes)
	err = service.User().Register(ctx, req.RegisterReq)
	res.Message = "注册成功"
	return
}

func (cUser) Info(ctx context.Context, req *user.InfoReq) (res *response.JsonRes, err error) {
	res = new(response.JsonRes)
	res.Data, err = service.User().Info(ctx, req.Uid)
	return
}

func (cUser) Menu(ctx context.Context, req *user.MenuReq) (res *response.JsonRes, err error) {
	res = new(response.JsonRes)
	res.Data, err = service.User().Menu(ctx)

	return
}

func (cUser) AccessCodes(ctx context.Context, req *user.AccessCodeReq) (res *response.JsonRes, err error) {
	res = new(response.JsonRes)
	codes, err := service.User().AccessCode(ctx, req.Role)
	res.Data = g.Map{
		"codes": codes,
		"uid":   req.Uid,
	}
	return
}
