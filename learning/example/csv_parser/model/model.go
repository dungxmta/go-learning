package model

var Keys = [...]string{
	"id",
	"value",
	"owner",
	"groups",
}

type Row struct {
	Type  string // ip / cid / sid
	Value string
	Group string
	Owner string
}
