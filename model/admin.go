package model

import "time"

type Admin struct {
	AdminId    int64     `xorm:"pk autoincr" json:"id"`
	AdminName  string    `xorm:"varchar(32)" json:"admin_name"`
	CreateTime time.Time `xorm:"DateTime" json:"create_time"`
	Status     int64     `xorm:"default 0" json:"status"`
	Avatar     string    `xorm:"varchar(255)" json:"avatar"`
	Pwd        string    `xorm:"varchar(255)" json:"pwd"`
	CityName   string    `xorm:"varchar(12)" json:"city_name"`
	CityId     int64     `xorm:"index" json:"city_id"`
	//City *City `xorm:"- <- ->"`
}
