package orm

import "errors"

var (
	// 存在审批中的变更请求
	ErrExistApprovingChangeRequest = errors.New("存在审批中的变更请求")
	// 变更请求不存在或已审批
	ErrChangeRequestNotExistOrApproved = errors.New("变更请求不存在或已审批")
)
