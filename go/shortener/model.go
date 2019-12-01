package shortener

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Url struct {
	ID        int64      `json:"id"`
	Url       string     `json:"url"`
	Shortid   string     `json:"shortId"`
	CreatedAt *time.Time `json:"createdAt" gorm:"default:CURRENT_DATE"`
	UpdatedAt *time.Time `json:"updatedAt" gorm:"default:null"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"default:null"`
}

func (Url) TableName() string {
	return "shortened_urls"
}

func (url *Url) findById(db *gorm.DB) error {
	return db.First(url, url.ID).Error
}

func (url *Url) findByUrl(db *gorm.DB) error {
	return db.First(url, "url = ?", url.Url).Error
}

func (url *Url) findByShortid(db *gorm.DB) error {
	return db.First(url, "shortid = ?", url.Shortid).Error
}

func (url *Url) create(db *gorm.DB) error {
	return db.Create(url).Error
}
