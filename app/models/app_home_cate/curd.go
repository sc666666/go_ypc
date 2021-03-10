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
func GetInfo(size int, page int, title string, subtitle string) (map[string]interface{}, error) {
	var appHomeCate []AppHomeCate
	//var cateData AppHomeCate
	var count int64
	db := model.DB

	fmt.Print(title)
	//可以绑定参数不需要传参搜索
	//if cateData.Title != "" {
	if title != "" {
		db = db.Where("title = ?", title)
	}
	if subtitle != "" {
		db = db.Where("subtitle = ?", subtitle)
	}

	db.Find(&appHomeCate).Count(&count) //总行数
	//model.DB.Find(&appHomeCate).Count(&count) //总行数
	fmt.Print(count)
	//不带条件分页查询
	if err := db.Offset((page-1)*size).Limit(size).Find(&appHomeCate).Error; err != nil {
	//if err := model.DB.Offset((page-1)*size).Limit(size).Find(&appHomeCate).Error; err != nil {
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
		"msg": "创建成功",
	} , nil
}

// 更新数据
func SaveInfo(appHomeCate *AppHomeCate, id interface{}) (map[string]interface{}, error) {
	if err := model.DB.First(&appHomeCate, id).Error; err != nil {
	//if err := model.DB.Where("id = ?", id).First(&appHomeCate).Error; err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		if err := model.DB.Save(&appHomeCate).Error; err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"data": appHomeCate,
			"msg": "更新成功",
		} , nil
	}
}

// 删除数据
func DeleteById(id interface{}) (map[string]interface{}, error) {
	var appHomeCate AppHomeCate
	if err := model.DB.First(&appHomeCate, id).Error; err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		if err := model.DB.Delete(&appHomeCate, id).Error; err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"data": appHomeCate,
			"msg": "删除成功",
		} , nil
	}
}
