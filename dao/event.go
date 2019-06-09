package dao

type Event struct {
	ID          int64  `gorm:"primary_key";AUTO_INCREMENT`
	Title       string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:varchar(1000)"`
	StartDate   int64
	EndDate     int64
	CategoryID  int64
}
