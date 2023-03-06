package models

type RentMovie struct {
	ID      uint `gorm:"primaryKey"`
	MovieID uint
	Movie   Movie `gorm:"foreignKey:MovieID"`
	RentID  uint
	Rent    Rent `gorm:"foreignKey:RentID"`
}
