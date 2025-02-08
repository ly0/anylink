package dbdata

import (
	"errors"
	"time"
)

type IpMap struct {
	Id        int       `json:"id" xorm:"pk autoincr not null"`
	IpAddr    string    `json:"ip_addr" xorm:"varchar(32) not null unique"`
	Username  string    `json:"username" xorm:"varchar(60) not null"`
	MacAddr   string    `json:"mac_addr" xorm:"varchar(32)"` // MAC地址
	UniqueMac bool      `json:"unique_mac" xorm:"Bool"` // MAC地址唯一性标志
	Keep      bool      `json:"keep" xorm:"Bool"` // 保留 IP 绑定
	KeepTime  time.Time `json:"keep_time" xorm:"DateTime"`
	Note      string    `json:"note" xorm:"varchar(255)"` // 备注
	LastLogin time.Time `json:"last_login" xorm:"DateTime"`
	UpdatedAt time.Time `json:"updated_at" xorm:"DateTime updated"`
}

func SetIpMap(v *IpMap) error {
	var err error

	if len(v.IpAddr) < 4 {
		return errors.New("IP地址错误")
	}

	if len(v.Username) == 0 {
		return errors.New("用户名不能为空")
	}

	v.UpdatedAt = time.Now()
	if v.Id > 0 {
		err = Set(v)
	} else {
		err = Add(v)
	}
	return err
}
