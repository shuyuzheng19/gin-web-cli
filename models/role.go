package models

import (
	"gin-web/common"
)

// Role 角色模型
type Role struct {
	ID          uint   `gorm:"primary_key;type:int;comment:角色ID"`
	Name        string `gorm:"size:255;unique;not null;comment:角色名"`
	Description string `gorm:"size:255;not null;comment:角色描述"`
}

func (*Role) TableName() string { return common.RoleTableName }

func GetDefaultRoles() []Role {
	return []Role{
		{ID: uint(common.USER_ID), Name: "USER", Description: "普通用户"},
		{ID: uint(common.ADMIN_ID), Name: "ADMIN", Description: "管理员"},
		{ID: uint(common.SUPER_ADMIN_ID), Name: "SUPER_ADMIN", Description: "超级管理员"},
	}
}
