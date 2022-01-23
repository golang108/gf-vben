// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dto

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Router is the golang structure of table router for DAO operations like Where/Data.
type Router struct {
	g.Meta    `orm:"table:router, dto:true"`
	Id        interface{} //
	Path      interface{} //
	Name      interface{} //
	Redirect  interface{} //
	Title     interface{} //
	Icon      interface{} //
	Component interface{} //
	Parent    interface{} //
	OrderNo   interface{} //
	Status    interface{} //
	CreateAt  *gtime.Time //
	UpdateAt  *gtime.Time //
	DeleteAt  *gtime.Time //
}
