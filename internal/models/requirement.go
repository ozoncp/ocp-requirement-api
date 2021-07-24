package models

import "fmt"

type Requirement struct {
	Id     uint64
	UserId uint64
	Text   string
}

func (r Requirement) String() string {
	return fmt.Sprintf(
		"{Id:%d, UserId:%d, Text:\"%s\"}",
		r.Id,
		r.UserId,
		r.Text)
}
