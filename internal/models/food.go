package models

type Food struct {
	ID          uint     `json:"id" gorm:"primaryKey"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Cost        int      `json:"cost"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `json:"category" gorm:"foreignKey:CategoryID;association_autoupdate:false;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Picture     string   `json:"picture"`
}

func (f Food) TableName() string {
	return "food"
}

func NewFood(name string, description string, cost int, categoryID uint, picture string) Food {
	return Food{
		Name:        name,
		Description: description,
		Cost:        cost,
		CategoryID:  categoryID,
		Picture:     picture,
	}
}
