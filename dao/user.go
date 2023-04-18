package dao

import (
	"gin-blog/database"
	"gin-blog/models"
)

func CreatedOrUpdate(userinfo *models.User) error {
	var sqluserinfo *models.User
	results := database.DB.Where("social_source= ? and social_user_id = ?", userinfo.SocialSource, userinfo.SocialUserId).First(&sqluserinfo)
	if results.Error != nil {
		results := database.DB.Create(&userinfo)
		if results.Error != nil {
			return results.Error
		}

	} else {
		userinfo.Id = sqluserinfo.Id
		results := database.DB.Model(&userinfo).Updates(&userinfo)
		return results.Error
	}
	return nil
}

// 后台登录
func GetUser(u, p string) (models.User, error) {
	var userinfo models.User
	err := database.DB.Table("users").Where("username = ? and password = ? ", u, p).First(&userinfo).Error
	if err != nil {
		return userinfo, err
	}
	return userinfo, nil
}

//查看所有用户

func GetUserList(ps, pn int) ([]models.User, int64, error) {
	var total int64
	list := make([]models.User, 0)
	err := database.DB.
		Table("users").
		Count(&total).
		Limit(pn).
		Offset(pn * (ps - 1)).
		Find(&list).Error
	if err != nil {
		return list, total, err
	}
	return list, total, nil
}

//查看某个用户 根据id

func GetUserById(i int) (models.User, error) {
	var userinfo models.User
	err := database.DB.Where("id = ?", i).Find(&userinfo).Error
	if err != nil {
		return userinfo, err
	}
	return userinfo, nil
}
