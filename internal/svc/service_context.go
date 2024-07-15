package svc

import (
	"github.com/WTFAcademy/platform/internal/config"
	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	DB     sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     sqlx.NewSqlConn(c.Database.Type, c.Database.Source),
	}
}
