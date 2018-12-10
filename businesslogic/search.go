package businesslogic

import "github.com/jinzhu/gorm"

// Search model for database
type SearchHistory struct {
	gorm.Model
	AccountID int    `gorm:"NOT NULL;REFERENCES ACCOUNTS(ID)"`
	Query     string `gorm:"TYPE:TEXT;NOT NULL"`
}
