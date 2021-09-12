package errcode

import "github.com/lynndotconfig/gkgo2/pkg/gkgo2err"

var (
	SystemError = gkgo2err.Common().New(gkgo2err.System)
	SystemDBError = SystemError.New(1, "init db error")
	
	UserError = gkgo2err.Business().New(gkgo2err.User)
	UserEmailNotFound = UserError.New(1, "not found user by email")
)
