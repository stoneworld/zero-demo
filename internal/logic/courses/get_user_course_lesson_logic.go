package courses

import (
	"context"
	"fmt"

	"github.com/WTFAcademy/platform/internal/svc"
	"github.com/WTFAcademy/platform/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserCourseLessonLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserCourseLessonLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserCourseLessonLogic {
	return &GetUserCourseLessonLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserCourseLessonLogic) GetUserCourseLesson(req *types.GetUserCourseLessonReq) error {
	// todo: add your logic here and delete this line
	fmt.Println(req.CourseId)
	return nil
}
