package courses

import (
	"net/http"

	"github.com/WTFAcademy/platform/internal/logic/courses"
	"github.com/WTFAcademy/platform/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AllCourseHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := courses.NewAllCourseLogic(r.Context(), svcCtx)
		resp, err := l.AllCourse()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
