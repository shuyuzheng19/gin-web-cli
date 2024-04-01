package user

import (
	"fmt"
	"gin-web/common"
	"gin-web/models"
	"gin-web/request"
	"gin-web/response"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	table string
	db    *gorm.DB
}

func (u UserRepositoryImpl) Create(user models.User) error {
	return u.db.Table(u.table).Create(&user).Error
}

func (u UserRepositoryImpl) GetById(id int) *models.User {
	var user models.User
	if err := u.db.Table(u.table).Preload("Role").First(&user, "id = ?", id); err != nil {
		return nil
	}
	return &user
}

func (u UserRepositoryImpl) FindByUsername(username string) *models.User {
	var user models.User
	if err := u.db.Table(u.table).Preload("Role").First(&user, "username = ?", username).Error; err != nil {
		return nil
	}
	return &user
}

func (u UserRepositoryImpl) FindByUsernameAndEmail(username, email string) *models.User {
	var user models.User
	if err := u.db.Table(u.table).First(&user, "email = ? and username = ?", email, username).Error; err != nil {
		return nil
	}
	return &user
}

func (u UserRepositoryImpl) GetUserByPage(req request.UserFilter) (users []response.UserResponse, total int64) {

	var pageCount = common.UserPageSize

	fmt.Println(req.Page, req.RoleID)

	var build = u.db.Model(&models.User{}).Table(u.table + " u")

	if req.RoleID > 0 {
		build.Where("u.role_id = ?", req.RoleID)
	}

	if build.Count(&total); total == 0 {
		return make([]response.UserResponse, 0), 0
	}

	build.Joins(fmt.Sprintf("join %s r on u.role_id = r.id", common.RoleTableName)).
		Select("u.id,u.nick_name as nickname,u.username,u.avatar,r.id as role_id").Scan(&users).
		Offset((req.Page - 1) * pageCount).
		Limit(pageCount).
		Order(req.Sort.GetOrderString("u.")).Scan(&users)

	return users, total
}

func (u UserRepositoryImpl) UpdatePassword(username, email, newPassword string) int64 {
	return u.db.Table(u.table).Where("username = ? and email = ?", username, email).UpdateColumn("password", newPassword).RowsAffected
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepositoryImpl{
		table: common.UserTableName,
		db:    db,
	}
}
