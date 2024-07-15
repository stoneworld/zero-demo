package courses

import (
	"net/http"

	"github.com/WTFAcademy/platform/internal/logic/courses"
	"github.com/WTFAcademy/platform/internal/svc"
	"github.com/WTFAcademy/platform/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetUserCourseLessonHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetUserCourseLessonReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := courses.NewGetUserCourseLessonLogic(r.Context(), svcCtx)
		err := l.GetUserCourseLesson(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
