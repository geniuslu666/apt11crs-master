package input_app_member

type FxMemberInfo struct {
	AuthId  string `json:"auth_id"    v:"required#Please select the  auth_id" dc:"用户授权编号"`
	Channel string `json:"channel"   v:"required#Please select the channel" dc:"渠道"`
}
