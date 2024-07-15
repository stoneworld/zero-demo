package courses

import (
	"net/http"

	"github.com/WTFAcademy/platform/internal/logic/courses"
	"github.com/WTFAcademy/platform/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCourseInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := courses.NewGetCourseInfoLogic(r.Context(), svcCtx)
		err := l.GetCourseInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
