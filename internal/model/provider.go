package model

type Provider struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"column:name;unique"`
	APIURL       string `gorm:"column:api_url"`
	ResponseKeys string `gorm:"column:response_keys"`
}
