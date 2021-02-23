package app_home_cate

import "go_ypc/pkg/model"

// 通过 ID 查询
func GetInfoById(id interface{}) (AppHomeCate, error) {
	var appHomeCate AppHomeCate

	if err := model.DB.First(&appHomeCate, id).Error; err != nil {
		return appHomeCate, err
	}

	return appHomeCate, nil
}

// 查询所有数据
//func (d *Dao) QueryHuman() (list []*model.HumanStats, err error) {
//	if err = d.db.Find(&list); err != nil {
//		return
//	}
//	return
//}

func GetInfo() (AppHomeCate, error) {
	var appHomeCate AppHomeCate

	if err := model.DB.Find(&appHomeCate).Error; err != nil {
		return appHomeCate, err
	}

	return appHomeCate, nil
}

