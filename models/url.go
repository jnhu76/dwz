package models

import (
	"gorm.io/gorm"
)

type Url struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	OriginUrl   string `json:"origin_url" gorm:"index"`
	ShorternUrl string `json:"shortern_url"`
	CreatedBy   int    `json:"created_by"`

	Model
}

func ExistUrlByShortern(shorten string) (bool, error) {
	var url Url
	err := db.Where(&Url{ShorternUrl: shorten}).First(&url).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, nil
	}

	if url.ID > 0 {
		return true, nil
	}

	return false, nil
}

// ExistUrlByOrigin checks if a url exists based on origin url
func ExistUrlByOrigin(origin string) (bool, error) {
	var url Url

	err := db.Where(&Url{OriginUrl: origin}).First(&url).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if url.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetUrlsByCreator get a list of urls based on creator id
func getUrlsByCreator(creator int) ([]*Url, error) {
	var urls []*Url

	err := db.Where(&Url{CreatedBy: creator}).Find(&urls).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return urls, nil
}

// GetUrl get a url based on id
func GetUrl(id int) (*Url, error) {
	var url Url
	err := db.Where(&Url{ID: id}).First(&url).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &url, nil
}

// GetUrl get a url based on short_url
func GetUrlByShort(short_url string) (*Url, error) {
	var url Url
	err := db.Where(&Url{ShorternUrl: short_url}).First(&url).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &url, nil
}

// GetShortByOrigin get a shortern url on origin_url
func GetShortByOrigin(origin_url string) (*Url, error) {
	var url Url
	err := db.Where(&Url{OriginUrl: origin_url}).First(&url).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &url, nil
}

// GetUrls get urls based on creator
func GetUrls(creator int) ([]*Url, error) {
	var urls []*Url
	err := db.Where(&Url{CreatedBy: creator}).Find(&urls).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return urls, nil
}

// DeleteUrl delete a single url
func DeleteUrl(shortern string) error {
	if err := db.Where("shortern_url = ?", shortern).Delete(&Url{}).Error; err != nil {
		return err
	}
	return nil
}

// AddUrl add a single url
func AddUrl(data map[string]interface{}) error {
	url := Url{
		OriginUrl:   data["origin_url"].(string),
		ShorternUrl: data["shorter_url"].(string),
		CreatedBy:   data["created_by"].(int),
	}

	if err := db.Create(&url).Error; err != nil {
		return err
	}
	return nil
}
