package course

import (
	"context"

	"github.com/WTFAcademy/platform/internal/model"
	"github.com/WTFAcademy/platform/internal/types"
)

type CourseBiz struct {
	ModelCourse model.TbCourseModel
}

func NewCourseBiz(modelCourse model.TbCourseModel) *CourseBiz {
	return &CourseBiz{
		ModelCourse: modelCourse,
	}
}

func (biz *CourseBiz) GetCourseList(ctx context.Context, status int) (list []types.Course) {
	list = make([]types.Course, 0)
	if courseList, err := biz.ModelCourse.FindOnlineCourses(status); err != nil {
		return list
	} else {
		for _, course := range courseList {
			list = append(list, types.Course{
				ID:          course.Id,
				Title:       course.Title,
				Description: course.Description,
				CoverImg:    course.CoverImg,
				RoutePath:   course.RoutePath,
				TotalScore:  course.TotalScore,
				UserCnt:     course.UserCnt,
				ShareURL:    course.ShareUrl,
				StartStatus: int(course.StartStatus),
			})
		}
	}
	return list
}
