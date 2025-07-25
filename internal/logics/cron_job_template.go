package logics

import (
	v1 "DeviceManagement/api/v1"
	"DeviceManagement/internal/dao"
	"DeviceManagement/internal/model"
	"DeviceManagement/internal/model/entity"
	"context"
	"database/sql"
	"sync"
)

var (
	cronJobTemplateOnce     sync.Once
	cronJobTemplateInstance *cronJobTemplate
)

type cronJobTemplate struct{}

func NewCronJobTemplate() *cronJobTemplate {
	cronJobTemplateOnce.Do(func() {
		cronJobTemplateInstance = &cronJobTemplate{}
	})
	return cronJobTemplateInstance
}

func (c *cronJobTemplate) RegisterDefaultCronJobTemplate() {
	for _, v := range model.DefaultCronJobList {
		dataInsert := map[string]interface{}{
			dao.CronJobTemplate.Columns().Name:       v.Name,
			dao.CronJobTemplate.Columns().InvokeType: v.InvokeType,
			dao.CronJobTemplate.Columns().Config:     v.Config,
		}

		exists, err := dao.CronJobTemplate.Ctx(context.Background()).Where(dao.CronJobTemplate.Columns().Name, v.Name).Exist()
		if err != nil {
			panic(err)
		}
		if exists {
			continue
		}
		_, err = dao.CronJobTemplate.Ctx(context.Background()).Data(dataInsert).Insert()
		if err != nil {
			panic(err)
		}
	}
}

func (c *cronJobTemplate) Get(ctx context.Context, id string) (out *model.CronJobTemplate, err error) {
	var entity *entity.TCronJobTemplate
	err = dao.CronJobTemplate.Ctx(ctx).Where(dao.CronJobTemplate.Columns().ID, id).Scan(&entity)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return
	}
	out = c.convertEntityToModel(entity)
	return
}

func (c *cronJobTemplate) List(ctx context.Context, in *v1.ListCronJobTemplateReq) (out []*model.CronJobTemplate, pageRes *model.PageRes, err error) {
	pageRes = &model.PageRes{}
	if in.Page == 0 {
		in.Page = 1
	}
	if in.PageSize == 0 {
		in.PageSize = model.DefaultPageSize
	}
	pageRes.CurrentPage = in.Page

	m := dao.CronJobTemplate.Ctx(ctx)

	pageRes.Total, err = m.Count()
	if err != nil {
		return
	}

	var entityList []*entity.TCronJobTemplate
	err = m.Page(in.Page, in.PageSize).Scan(&entityList)

	for _, v := range entityList {
		out = append(out, c.convertEntityToModel(v))
	}
	return
}

func (c *cronJobTemplate) convertEntityToModel(in *entity.TCronJobTemplate) (out *model.CronJobTemplate) {
	out = &model.CronJobTemplate{
		ID:         in.ID,
		Name:       in.Name,
		InvokeType: in.InvokeType,
		Config:     in.Config,
		CreatedAt:  in.CreatedAt,
		UpdatedAt:  in.UpdatedAt,
	}
	return
}
