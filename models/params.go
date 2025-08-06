package models

// 资料修改参数
type EditableProfileParams struct {
	Nickname string `json:"nickname" form:"nickname"`
	Avatar   string `json:"avatar" form:"avatar"`
	Gender   int    `json:"gender" form:"gender"`
}
