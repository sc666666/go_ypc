package app_home_cate

import (
	"go_ypc/app/models"
	"time"
)

type AppHomeCate struct {
	models.BaseModel

	Title     string `json:"title" form:"title"`
	Subtitle  string `json:"subtitle"`
	Cate      uint64 `json:"cate"`
	Sort      uint64 `json:"sort"`
	Status    uint64 `json:"status"`
	CreatedAt time.Time
	CreatedBy string `json:"createdBy"`
	UpdatedAt time.Time
	UpdatedBy string `json:"updatedBy"`
}

func (AppHomeCate) TableName() string {
	return "erp_ypc_app_home_cate"
}
