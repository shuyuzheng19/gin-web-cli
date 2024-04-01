package upload

import (
	"fmt"
	"gin-web/common"
	"gin-web/models"
	"gin-web/request"
	"gin-web/response"

	"gorm.io/gorm"
)

type FileRepositoryImpl struct {
	db        *gorm.DB
	tableName string
}

func (u FileRepositoryImpl) FindByMd5(md5 string) string {
	var r string
	u.db.Model(&models.FileMd5Info{}).Select("url").Where("md5=?", md5).Scan(&r)
	return r
}

func (u FileRepositoryImpl) BatchSave(files []models.FileInfo) error {
	return u.db.Model(&models.FileInfo{}).Save(&files).Error
}

func (u FileRepositoryImpl) FindFileInfos(uid int, req request.FileRequest) (_ []response.FileResponse, count int64) {

	var files = make([]response.FileResponse, 0)

	var build = u.db.Model(&models.FileInfo{}).Table(u.tableName + " f")

	if uid > 0 {
		build.Where("f.user_id = ?", uid)
	} else {
		build.Where("f.is_pub = ?", true)
	}

	if req.Keyword != "" {
		build.Where("old_name like ?", "%"+req.Keyword+"%")
	}

	if build.Count(&count); count == 0 {
		return files, 0
	}

	build.Joins(fmt.Sprintf("join %s fm on fm.md5 = f.md5", common.FileMd5TableName))

	var pageCount = common.FilePageCount

	build.Select("f.id,f.old_name as name,f.created_at,f.suffix,f.size",
		"fm.md5 as md5,fm.url as url").Offset((req.Page - 1) * pageCount).Limit(pageCount)

	if req.Sort == "size" {
		build.Order("f.size desc")
	} else {
		build.Order("f.created_at desc")
	}

	build.Scan(&files)

	return files, count
}

func NewFileRepository(db *gorm.DB) FileRepository {
	return FileRepositoryImpl{db: db, tableName: common.FileTableName}
}
