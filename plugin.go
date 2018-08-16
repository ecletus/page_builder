package page_builder

import (
	"github.com/moisespsena-go/aorm"
)

type Plugin struct {

}


func (Plugin) SetupSystemDB(db *aorm.DB) interface{} {
	return nil
}