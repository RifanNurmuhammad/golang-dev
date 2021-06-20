package domain

import "time"

type LogActivities struct {
	ID              int32     `gorm:"primary_key" json:"id" jsonapi:"primary,logActivities"`
	SearchCall      string    `gorm:"column:searchCall" json:"searchCall"`
	CreatedAt       time.Time `gorm:"column:createdAt" json:"createdAt"`
	CreatedAtString string    `gorm:"-" sql:"-" jsonapi:"attr,createdAt"`
}
