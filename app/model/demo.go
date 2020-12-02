package model

import (
	"strconv"
)

type Demo struct {
	Id   *int    `json:"id"`
	Name *string `json:"name"`
	Wxid *string `json:"wxid"`
}

func (u *Demo) String() string {
	return strconv.Itoa(*u.Id) + " " + *u.Name + " " + *u.Wxid
}

func (Demo) TableName() string {
	return "sys_user"
}
