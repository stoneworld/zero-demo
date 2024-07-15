package courses

import (
	"context"

	"github.com/WTFAcademy/platform/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCourseQuizListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCourseQuizListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCourseQuizListLogic {
	return &GetCourseQuizListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCourseQuizListLogic) GetCourseQuizList() error {
	// todo: add your logic here and delete this line

	return nil
}
