package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysUserModel = (*customSysUserModel)(nil)

type (
	// SysUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUserModel.
	SysUserModel interface {
		sysUserModel
		withSession(session sqlx.Session) SysUserModel
		FindAll(ctx context.Context, page, pageSize int, account string, status int) (*[]SysUser, int64, error)
	}

	customSysUserModel struct {
		*defaultSysUserModel
	}
)

// NewSysUserModel returns a model for the database table.
func NewSysUserModel(conn sqlx.SqlConn) SysUserModel {
	return &customSysUserModel{
		defaultSysUserModel: newSysUserModel(conn),
	}
}

func (m *customSysUserModel) withSession(session sqlx.Session) SysUserModel {
	return NewSysUserModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customSysUserModel) FindAll(ctx context.Context, page, pageSize int, account string, statusId int) (*[]SysUser, int64, error) {
	var resp []SysUser
	var count int64
	return &resp, count, nil
}
