package app_home_cate

import (
	"fmt"
	"go_ypc/pkg/model"
)

// 通过 ID 查询
func GetInfoById(id interface{}) (AppHomeCate, error) {
	var appHomeCate AppHomeCate

	if err := model.DB.First(&appHomeCate, id).Error; err != nil {
		return appHomeCate, err
	}

	return appHomeCate, nil
}

// 查询所有数据
func GetInfo(size int, page int) (map[string]interface{}, error) {
	var appHomeCate []AppHomeCate
	var count int64
	model.DB.Find(&appHomeCate).Count(&count) //总行数
	fmt.Print(count)
	if err := model.DB.Offset((page-1)*size).Limit(size).Find(&appHomeCate).Error; err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"homeCate":appHomeCate,
		"total":count,
	} , nil
}

// 创建数据
func CreateInfo(appHomeCate *AppHomeCate) (map[string]interface{}, error) {
	if err := model.DB.Create(&appHomeCate).Error; err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"data": appHomeCate,
	} , nil
}

// 更新数据
func SaveInfo(appHomeCate *AppHomeCate, id int) (map[string]interface{}, error) {


	if err := model.DB.Where("id = ?", id).First(&appHomeCate).Error; err != nil {
		return nil, err
	} else {
		if err := model.DB.Save(&appHomeCate).Error; err != nil {
			return nil, err
		}

		return map[string]interface{}{
			"data": appHomeCate,
		} , nil
	}


	//if err := model.DB.Save(&appHomeCate).Error; err != nil {
	//	return nil, err
	//}
	//return map[string]interface{}{
	//	"data": appHomeCate,
	//} , nil
}
