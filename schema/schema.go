package schema
import (
	"gorm.io/gorm"
)
type DBconnection struct{
	DB *gorm.DB
}
var DBConnect = &DBconnection{}
type Url_Data struct{
	gorm.Model
	Short_url string `gorm:"uniqueIndex"`
	Long_url string  `gorm:"uniqueIndex"`
}
