package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RbacUserModel = (*customRbacUserModel)(nil)

type (
	// RbacUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRbacUserModel.
	RbacUserModel interface {
		rbacUserModel
		withSession(session sqlx.Session) RbacUserModel
	}

	customRbacUserModel struct {
		*defaultRbacUserModel
	}
)

// NewRbacUserModel returns a model for the database table.
func NewRbacUserModel(conn sqlx.SqlConn) RbacUserModel {
	return &customRbacUserModel{
		defaultRbacUserModel: newRbacUserModel(conn),
	}
}

func (m *customRbacUserModel) withSession(session sqlx.Session) RbacUserModel {
	return NewRbacUserModel(sqlx.NewSqlConnFromSession(session))
}
