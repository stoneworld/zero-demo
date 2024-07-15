package courses

import (
	"context"

	"github.com/WTFAcademy/platform/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetCourseInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCourseInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCourseInfoLogic {
	return &GetCourseInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCourseInfoLogic) GetCourseInfo() error {
	// todo: add your logic here and delete this line

	return nil
}
