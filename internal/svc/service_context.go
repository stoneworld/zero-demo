package svc

import (
	"github.com/WTFAcademy/platform/internal/biz/course"
	"github.com/WTFAcademy/platform/internal/config"
	"github.com/WTFAcademy/platform/internal/model"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config      config.Config
	DB          sqlx.SqlConn
	ModelCourse *model.TbCourseModel
	BizCourse   *course.CourseBiz
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := sqlx.NewSqlConn(c.Database.Type, c.Database.Source)
	modelCourse := model.NewTbCourseModel(db)
	bizCourse := course.NewCourseBiz(modelCourse)
	return &ServiceContext{
		Config:      c,
		DB:          db,
		ModelCourse: &modelCourse,
		BizCourse:   bizCourse,
	}
}
