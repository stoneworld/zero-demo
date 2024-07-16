package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TbUserCoreModel = (*customTbUserCoreModel)(nil)

type (
	// TbUserCoreModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbUserCoreModel.
	TbUserCoreModel interface {
		tbUserCoreModel
		withSession(session sqlx.Session) TbUserCoreModel
	}

	customTbUserCoreModel struct {
		*defaultTbUserCoreModel
	}
)

// NewTbUserCoreModel returns a model for the database table.
func NewTbUserCoreModel(conn sqlx.SqlConn) TbUserCoreModel {
	return &customTbUserCoreModel{
		defaultTbUserCoreModel: newTbUserCoreModel(conn),
	}
}

func (m *customTbUserCoreModel) withSession(session sqlx.Session) TbUserCoreModel {
	return NewTbUserCoreModel(sqlx.NewSqlConnFromSession(session))
}
