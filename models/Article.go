package models

import (
	"ginblog/utils/errmsg"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Category     Category `gorm:"foreignkey:Cid"`
	Title        string   `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int      `gorm:"type:int;not null" json:"cid"`
	Desc         string   `gorm:"type:varchar(200)" json:"desc"`
	Content      string   `gorm:"type:longtext" json:"content"`
	Img          string   `gorm:"type:varchar(200)" json:"img"`
	CommentCount int      `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int      `gorm:"type:int;not null;default:0" json:"read_count"`
}

// 新增文章
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetArticle 查询文章列表
func GetArticle(pageSize int, pageNum int) ([]Article, int, int) {
	var articleList []Article
	var err error
	var total int64

	err = db.Select("articles.id, title, img, created_at, updated_at, `desc`, comment_count, read_count, Category.name").Limit(pageSize).Offset((pageNum - 1) * pageSize).Order("created_at DESC").Joins("Category").Find(&articleList).Error
	// 单独计数
	db.Model(&articleList).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR, 0
	}
	return articleList, errmsg.SUCCESS, int(total)

}

// SearchArticle 搜索文章标题
func SearchArticle(cid int, title string, pageSize int, pageNum int) ([]Article, int, int64) {
	var articleList []Article
	var err error
	var total int64
	if cid == 0 {
		err = db.Select("articles.id,title, img, created_at, updated_at, `desc`, comment_count, read_count, Category.name").Order("Created_At DESC").Joins("Category").Where("title LIKE ?",
			title+"%",
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
		//单独计数
		db.Model(&articleList).Where("title LIKE ?",
			title+"%",
		).Count(&total)
		if err != nil {
			return nil, errmsg.ERROR, 0
		}
	} else {
		err = db.Select("articles.id,title, img, created_at, updated_at, `desc`, comment_count, read_count, Category.name").Order("Created_At DESC").Joins("Category").Where("articles.cid = ? and title LIKE ?", cid,
			title+"%",
		).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
		//单独计数
		db.Model(&articleList).Where("articles.cid = ? and title LIKE ?", cid,
			title+"%",
		).Count(&total)
		if err != nil {
			return nil, errmsg.ERROR, 0
		}
	}

	return articleList, errmsg.SUCCESS, total
}

// 查询分类下的所有文章
func GetArticleByCate(id int, pageSize int, pageNum int) ([]Article, int, int) {
	var cateArtlist []Article
	var total int64
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).Find(&cateArtlist).Error
	db.Model(&Article{}).Where("cid = ?", id).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtlist, errmsg.SUCCESS, int(total)
}

// 查询单个文章
func GetArticleInfo(id int) (*Article, int) {
	var art Article
	err := db.Preload("Category").Where("id = ?", id).First(&art).Error
	db.Model(&art).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
	if err != nil {
		return &art, errmsg.ERROR_ART_NOT_EXIST
	}
	return &art, errmsg.SUCCESS
}

// 编辑文章信息
func EditArticle(id int, data *Article) int {
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err := db.Model(&Article{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除文章
func DeleteArticle(id int) int {
	var article Article
	err := db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
