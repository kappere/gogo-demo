package model

import (
	"strconv"
)

type Demo struct {
	Id   *int    `json:"id"`
	Name *string `json:"name"`
	Wxid *string `json:"wxid"`
}

func (Demo) TableName() string {
	return "demo"
}

func (u *Demo) String() string {
	return strconv.Itoa(*u.Id) + " " + *u.Name + " " + *u.Wxid
}
