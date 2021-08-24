package models

import "fmt"

type Requirement struct {
	Id     uint64 `db:"id"`
	UserId uint64 `db:"user_id"`
	Text   string `db:"text"`
}

func (r Requirement) String() string {
	return fmt.Sprintf(
		"{Id:%d, UserId:%d, Text:\"%s\"}",
		r.Id,
		r.UserId,
		r.Text)
}
