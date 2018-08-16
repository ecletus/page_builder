package page_builder

import (
	"github.com/jinzhu/gorm"
)

type Plugin struct {

}


func (Plugin) SetupSystemDB(db *gorm.DB) interface{} {
	return nil
}