package model

import "time"

type Project struct {
	Name     string    `bson:"name"`
	Code     int       `bson:"code"`
	Sp       int       `bson:"sp"`
	Deadline time.Time `bson:"deadline"`
}
