package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// dicItem stores dictionary item
type DicItem struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Code        string        `json:"code"`
	Value       string        `json:"val"`
	Description string        `json:"desc"`
	CreatedAt   time.Time     `json:"crAt"`
	UpdatedAt   time.Time     `json:"upAt"`
}

// dicItem stores dictionary item
type MetaInfoItem struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Code        string        `json:"code"`
	Name        string        `json:"name"`
	Description string        `json:"desc"`
	CreatedAt   time.Time     `json:"crAt"`
	UpdatedAt   time.Time     `json:"upAt"`
	Category    string        `json:"cat"`
}
