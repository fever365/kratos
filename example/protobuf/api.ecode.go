// Code generated by protoc-gen-ecode v0.1, DO NOT EDIT.
// source: api.proto

package api

import (
<<<<<<< HEAD
	"github.com/fever365/kratos/pkg/ecode"
=======
	"github.com/fever365/kratos/pkg/ecode"
>>>>>>> 3c6dbc7bf446fcf807931c0adeb03ddb0e59f774
)

// to suppressed 'imported but not used warning'
var _ ecode.Codes

// UserErrCode ecode
var (
	UserNotExist         = ecode.New(-404)
	UserUpdateNameFailed = ecode.New(10000)
)
