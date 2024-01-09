package model

import "time"

type Task struct {
	RelatedProject Project   `bson:"relatedProject"`
	Executor       Executor  `bson:"executor"`
	IssueDate      time.Time `bson:"issueDate"`
	Sp             int       `bson:"sp"`
	ExpDeadline    time.Time `bson:"expDeadline"`
	RealDeadline   time.Time `bson:"realDeadline"`
}
