package courses

import (
	"context"

	"github.com/WTFAcademy/platform/internal/svc"
	"github.com/WTFAcademy/platform/internal/types"

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

func (l *AllCourseLogic) AllCourse() (resp *types.CourseListResp, err error) {
	courseList := l.svcCtx.BizCourse.GetCourseList(l.ctx, 1)
	return &types.CourseListResp{
		Data: types.CourseList{
			List: courseList,
		},
	}, nil
}
