package model

import (
	_ "github.com/lib/pq" // <------------ here
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"testing"
)

func Test_customTbCourseModel_FindOnlineCourses(t *testing.T) {
	conn := sqlx.NewSqlConn(`postgres`, `postgres://username:password@host:5432/postgres?sslmode=disable`)
	userModel := NewTbCourseModel(conn)
	got, err := userModel.FindOnlineCourses(1)
	if err != nil {
		t.Errorf("FindOnlineCourses() error = %v", err)
		return
	}
	t.Logf("FindOnlineCourses() got = %v", got)
}
