package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RbacRoleModel = (*customRbacRoleModel)(nil)

type (
	// RbacRoleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRbacRoleModel.
	RbacRoleModel interface {
		rbacRoleModel
		withSession(session sqlx.Session) RbacRoleModel
	}

	customRbacRoleModel struct {
		*defaultRbacRoleModel
	}
)

// NewRbacRoleModel returns a model for the database table.
func NewRbacRoleModel(conn sqlx.SqlConn) RbacRoleModel {
	return &customRbacRoleModel{
		defaultRbacRoleModel: newRbacRoleModel(conn),
	}
}

func (m *customRbacRoleModel) withSession(session sqlx.Session) RbacRoleModel {
	return NewRbacRoleModel(sqlx.NewSqlConnFromSession(session))
}
