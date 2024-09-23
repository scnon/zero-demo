package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysUserModel = (*customSysUserModel)(nil)

type (
	// SysUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUserModel.
	SysUserModel interface {
		sysUserModel
		withSession(session sqlx.Session) SysUserModel
		FindAll(ctx context.Context, page, pageSize int, account string, status int) (*[]SysUser, int, error)
		DeleteAll(ctx context.Context, ids []int64) error
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

func (m *customSysUserModel) FindAll(ctx context.Context, page, pageSize int, account string, status int) (*[]SysUser, int, error) {
	var count int
	query := fmt.Sprintf("select count(*) from %s where status = ?", m.table)
	err := m.conn.QueryRowCtx(ctx, &count, query, status)
	if err != nil {
		return nil, 0, err
	}

	resp := []SysUser{}
	if count == 0 {
		return &resp, count, nil
	}

	if account != "" {
		query = fmt.Sprintf("select * from %s where `account` = ? and `status` = ?", m.table)
		m.conn.QueryRowsCtx(ctx, &resp, query, account, status)
		return &resp, len(resp), nil
	}

	offset := (page - 1) * pageSize
	query = fmt.Sprintf("select * from %s where status = %d limit %d offset %d", m.table, status, pageSize, offset)
	m.conn.QueryRowsCtx(ctx, &resp, query)
	return &resp, count, nil
}

func (m *customSysUserModel) DeleteAll(ctx context.Context, ids []int64) error {
	_, err := m.conn.ExecCtx(ctx, "delete from sys_user where id in ?", ids)
	return err
}
