package admin

import (
	"time"
)

//AdminUser实体类
type XAdminRoleLimit struct {
	Limit time.Time `json:"limit"`
	Role  int       `json:"role"`
}

func (x *XAdminRoleLimit) IsEffectiveTime() bool {
	return time.Now().Before(x.Limit)
}
