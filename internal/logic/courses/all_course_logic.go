package courses

import (
	"context"

	"github.com/WTFAcademy/platform/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type AllCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllCourseLogic {
	return &AllCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllCourseLogic) AllCourse() error {
	// todo: add your logic here and delete this line

	return nil
}
