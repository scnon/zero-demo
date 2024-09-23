// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.2

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	rbacUserFieldNames          = builder.RawFieldNames(&RbacUser{})
	rbacUserRows                = strings.Join(rbacUserFieldNames, ",")
	rbacUserRowsExpectAutoSet   = strings.Join(stringx.Remove(rbacUserFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	rbacUserRowsWithPlaceHolder = strings.Join(stringx.Remove(rbacUserFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	rbacUserModel interface {
		Insert(ctx context.Context, data *RbacUser) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*RbacUser, error)
		FindOneByAccount(ctx context.Context, account string) (*RbacUser, error)
		Update(ctx context.Context, data *RbacUser) error
		Delete(ctx context.Context, id int64) error
	}

	defaultRbacUserModel struct {
		conn  sqlx.SqlConn
		table string
	}

	RbacUser struct {
		Id         int64          `db:"id"`          // 主键
		Account    string         `db:"account"`     // 账号
		UserName   string         `db:"user_name"`   // 昵称
		Password   string         `db:"password"`    // 密码
		Roles      sql.NullString `db:"roles"`       // 角色
		CreateTime time.Time      `db:"create_time"` // 创建时间
		UpdateTime sql.NullTime   `db:"update_time"` // 修改时间
	}
)

func newRbacUserModel(conn sqlx.SqlConn) *defaultRbacUserModel {
	return &defaultRbacUserModel{
		conn:  conn,
		table: "`rbac_user`",
	}
}

func (m *defaultRbacUserModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultRbacUserModel) FindOne(ctx context.Context, id int64) (*RbacUser, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", rbacUserRows, m.table)
	var resp RbacUser
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRbacUserModel) FindOneByAccount(ctx context.Context, account string) (*RbacUser, error) {
	var resp RbacUser
	query := fmt.Sprintf("select %s from %s where `account` = ? limit 1", rbacUserRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, account)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRbacUserModel) Insert(ctx context.Context, data *RbacUser) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, rbacUserRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Account, data.UserName, data.Password, data.Roles)
	return ret, err
}

func (m *defaultRbacUserModel) Update(ctx context.Context, newData *RbacUser) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, rbacUserRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, newData.Account, newData.UserName, newData.Password, newData.Roles, newData.Id)
	return err
}

func (m *defaultRbacUserModel) tableName() string {
	return m.table
}
