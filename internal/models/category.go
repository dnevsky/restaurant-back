package models

type Category struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Food []Food `json:"food,omitempty" gorm:"foreignKey:CategoryID"`
}

func (Category) TableName() string {
	return "categories"
}

func NewCategory(name string) Category {
	return Category{
		Name: name,
	}
}
