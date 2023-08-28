package domain

type Segment struct {
	ID   int    `json:"id" gorm:"primary_key" swaggerignore:"true"`
	Name string `json:"name" gorm:"not null;unique"`
}
