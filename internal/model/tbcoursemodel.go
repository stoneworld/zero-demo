package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TbCourseModel = (*customTbCourseModel)(nil)

type (
	// TbCourseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTbCourseModel.
	TbCourseModel interface {
		tbCourseModel
		withSession(session sqlx.Session) TbCourseModel
	}

	customTbCourseModel struct {
		*defaultTbCourseModel
	}
)

// NewTbCourseModel returns a model for the database table.
func NewTbCourseModel(conn sqlx.SqlConn) TbCourseModel {
	return &customTbCourseModel{
		defaultTbCourseModel: newTbCourseModel(conn),
	}
}

func (m *customTbCourseModel) withSession(session sqlx.Session) TbCourseModel {
	return NewTbCourseModel(sqlx.NewSqlConnFromSession(session))
}
