package model

type Executor struct {
	Id         int    `bson:"id"`
	FirstName  string `bson:"firstName"`
	LastName   string `bson:"lastName"`
	FatherName string `bson:"fatherName"`
	Position   string `bson:"position"`
}
