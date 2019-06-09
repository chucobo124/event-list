package dao

type Category struct {
	ID   int64 `gorm:"primary_key";AUTO_INCREMENT`
	Name string
}
