package response

import "cpm/model/system"

type SysUserResponse struct {
	User system.SysUser `json:"user"`
}
