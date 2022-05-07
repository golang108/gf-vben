// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Permission is the golang structure for table permission.
type Permission struct {
	Id       int64       `json:"id"       ` //
	Parent   int64       `json:"parent"   ` //
	Value    string      `json:"value"    ` //
	Desc     string      `json:"desc"     ` //
	CreateAt *gtime.Time `json:"createAt" ` //
	UpdateAt *gtime.Time `json:"updateAt" ` //
	DeleteAt *gtime.Time `json:"deleteAt" ` //
	Name     string      `json:"name"     ` //
	Type     int64       `json:"type"     ` // 1 权限域 2 权限组 3 权限
}
