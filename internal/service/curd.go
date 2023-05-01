// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"github.com/jinmao88/gf-utility/curd"

	"github.com/jinmao88/gf-utility/response"
)

type (
	ICurd interface {
		Curd(ctx context.Context, r *curd.CurdReq) (res *response.JsonRes, err error)
	}
)

var (
	localCurd ICurd
)

func Curd() ICurd {
	if localCurd == nil {
		panic("implement not found for interface ICurd, forgot register?")
	}
	return localCurd
}

func RegisterCurd(i ICurd) {
	localCurd = i
}
