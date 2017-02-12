package types

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//DicItem stores dictionary item
type DicItem struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Code      string        `json:"code"`
	Value     string        `json:"value"`
	Desc      string        `json:"desc"`
	DCode     string        `json:"dcode"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
}

// MetaInfoItem stores dictionary item
type MetaInfoItem struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Code        string        `json:"code"`
	Name        string        `json:"name"`
	Description string        `json:"desc"`
	CreatedAt   time.Time     `json:"crAt"`
	UpdatedAt   time.Time     `json:"upAt"`
	Category    string        `json:"cat"`
}
