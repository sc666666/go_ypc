package app_home_cate

import (
	"go_ypc/app/models"
	"time"
)

type AppHomeCate struct {
	models.BaseModel

	Title     string
	Subtitle  string
	Cate      uint64
	Sort      uint64
	Status    uint64
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}

func (AppHomeCate) TableName() string {
	return "erp_ypc_app_home_cate"
}
